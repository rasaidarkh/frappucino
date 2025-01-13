package handlers

import (
	"encoding/json"
	"fmt"
	"frappuccino/internal/models"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

type InventoryService interface {
	GetAll() ([]models.Inventory, error)
	GetElementById(id int) (models.Inventory, error)
	Delete(id int) error
	Put(item models.Inventory) error
	Post(item models.Inventory) error
}

type InventoryHandler struct {
	Service InventoryService
	Logger  *slog.Logger
}

func NewInventoryHandler(service InventoryService, logger *slog.Logger) *InventoryHandler {
	return &InventoryHandler{service, logger}
}

func (h *InventoryHandler) RegisterEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("POST /inventory", Middleware(h.Post))
	mux.HandleFunc("POST /inventory/", Middleware(h.Post))

	mux.HandleFunc("GET /inventory", Middleware(h.GetAll))
	mux.HandleFunc("GET /inventory/", Middleware(h.GetAll))

	mux.HandleFunc("GET /inventory/{id}", Middleware(h.GetElementById))
	mux.HandleFunc("GET /inventory/{id}/", Middleware(h.GetElementById))

	mux.HandleFunc("PUT /inventory/{id}", Middleware(h.Put))
	mux.HandleFunc("PUT /inventory/{id}/", Middleware(h.Put))

	mux.HandleFunc("DELETE /inventory/{id}", Middleware(h.Delete))
	mux.HandleFunc("DELETE /inventory/{id}/", Middleware(h.Delete))
}

func (h *InventoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	InventoryId, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Service.Delete(InventoryId)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode("deleted")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
	msg := fmt.Sprintf("entry with Id %v was deleted succefully\n", InventoryId)
	h.Logger.Info(msg)

}

func (h *InventoryHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *InventoryHandler) GetElementById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	InventoryId, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	inventoryItem, err := h.Service.GetElementById(InventoryId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(inventoryItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *InventoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	inventoryItems, err := h.Service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(inventoryItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h *InventoryHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error reading request body: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var singleItem models.Inventory
	if err := json.Unmarshal(data, &singleItem); err != nil {
		h.Logger.Error(fmt.Sprintf("error unmarshalling an inventory item: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Service.Post(singleItem); err != nil {
		h.Logger.Error(fmt.Sprintf("error creating single inventory item: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
