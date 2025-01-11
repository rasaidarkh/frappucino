package handlers

import (
	"encoding/json"
	"frappuccino/internal/service"
	"log/slog"
	"net/http"
)

type InventoryHandler struct {
	Service service.InventoryService
	Logger  *slog.Logger
}

func NewInventoryHandler(service service.InventoryService, logger *slog.Logger) *InventoryHandler {
	return &InventoryHandler{service, logger}
}

func (h *InventoryHandler) RegisterEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("POST /inventory", h.Post)
	mux.HandleFunc("POST /inventory/", h.Post)

	mux.HandleFunc("GET /inventory", h.GetAll)
	mux.HandleFunc("GET /inventory/", h.GetAll)

	mux.HandleFunc("GET /inventory/{id}", h.GetElementById)
	mux.HandleFunc("GET /inventory/{id}/", h.GetElementById)

	mux.HandleFunc("PUT /inventory/{id}", h.Put)
	mux.HandleFunc("PUT /inventory/{id}/", h.Put)

	mux.HandleFunc("DELETE /inventory/{id}", h.Delete)
	mux.HandleFunc("DELETE /inventory/{id}/", h.Delete)
}

func (h *InventoryHandler) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *InventoryHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *InventoryHandler) GetElementById(w http.ResponseWriter, r *http.Request) {

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

	w.WriteHeader(http.StatusOK)
}

func (h *InventoryHandler) Post(w http.ResponseWriter, r *http.Request) {

}
