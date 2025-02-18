package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"frappuccino/internal/handlers/middleware"
	"frappuccino/internal/models"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

type InventoryService interface {
	GetAll(ctx context.Context) ([]models.Inventory, error)
	GetElementById(ctx context.Context, id int) (models.Inventory, error)
	Delete(ctx context.Context, id int) error
	Put(ctx context.Context, item models.Inventory) error
	Post(ctx context.Context, item models.Inventory) error
}

type InventoryHandler struct {
	Service InventoryService
	Logger  *slog.Logger
}

func NewInventoryHandler(service InventoryService, logger *slog.Logger) *InventoryHandler {
	return &InventoryHandler{service, logger}
}

func (h *InventoryHandler) RegisterEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("POST /inventory", middleware.Middleware(h.Post))
	mux.HandleFunc("POST /inventory/", middleware.Middleware(h.Post))

	mux.HandleFunc("GET /inventory", middleware.Middleware(h.GetAll))
	mux.HandleFunc("GET /inventory/", middleware.Middleware(h.GetAll))

	mux.HandleFunc("GET /inventory/{id}", middleware.Middleware(h.GetElementById))
	mux.HandleFunc("GET /inventory/{id}/", middleware.Middleware(h.GetElementById))

	mux.HandleFunc("PUT /inventory/{id}", middleware.Middleware(h.Put))
	mux.HandleFunc("PUT /inventory/{id}/", middleware.Middleware(h.Put))

	mux.HandleFunc("DELETE /inventory/{id}", middleware.Middleware(h.Delete))
	mux.HandleFunc("DELETE /inventory/{id}/", middleware.Middleware(h.Delete))
}

func (h *InventoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	h.handleRequestWithID(w, r, func(ctx context.Context, id int) error {
		if err := h.Service.Delete(ctx, id); err != nil {
			h.Logger.Error(err.Error())
			return err
		}

		w.WriteHeader(204)
		h.Logger.Info("Inventory item was deleted", slog.Int("id", id))
		return nil
	})
}

func (h *InventoryHandler) Put(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error reading request body: %v", err))
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer r.Body.Close()

	var item models.Inventory
	if err := json.Unmarshal(data, &item); err != nil {
		h.Logger.Error(fmt.Sprintf("error unmarshalling inventory item: %v", err))
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := h.Service.Put(r.Context(), item); err != nil {
		h.Logger.Error(fmt.Sprintf("error updating inventory item: %v", err))
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to update item")
		return
	}

	w.WriteHeader(http.StatusOK)
	h.Logger.Info("Inventory item was updated", slog.Int("id", item.InventoryId))
}

func (h *InventoryHandler) GetElementById(w http.ResponseWriter, r *http.Request) {
	h.handleRequestWithID(w, r, func(ctx context.Context, id int) error {
		item, err := h.Service.GetElementById(ctx, id)
		if err != nil {
			h.Logger.Error(err.Error())
			return err
		}

		h.Logger.Info("Fetched an inventory item", slog.Int("id", id))
		return json.NewEncoder(w).Encode(item)
	})
}

func (h *InventoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.Service.GetAll(r.Context())
	if err != nil {
		h.Logger.Error(err.Error())
		h.writeErrorResponse(w, http.StatusBadRequest, "Failed to fetch inventory items")
		return
	}

	if err := json.NewEncoder(w).Encode(items); err != nil {
		h.Logger.Error(err.Error())
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}

	h.Logger.Info("Inventory items were fetched")
}

func (h *InventoryHandler) Post(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error reading request body: %v", err))
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer r.Body.Close()

	var item models.Inventory
	if err := json.Unmarshal(data, &item); err != nil {
		h.Logger.Error(fmt.Sprintf("error unmarshalling inventory item: %v", err))
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := h.Service.Post(r.Context(), item); err != nil {
		h.Logger.Error(fmt.Sprintf("error creating inventory item: %v", err))
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to create item")
		return
	}

	w.WriteHeader(http.StatusCreated)
	h.Logger.Info("New inventory item was added", slog.Int("id", item.InventoryId))
}

func (h *InventoryHandler) handleRequestWithID(w http.ResponseWriter, r *http.Request, handler func(context.Context, int) error) {
	idStr := r.URL.Path[len(r.URL.Path)-1:]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("invalid id: %v", err))
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := handler(r.Context(), id); err != nil {
		h.Logger.Error(err.Error())
		h.writeErrorResponse(w, http.StatusBadRequest, "No such inventory item")
	}
}

func (h *InventoryHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}
