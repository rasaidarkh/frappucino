package handlers

import (
	"context"
	"frappuccino/internal/handlers/middleware"
	"log/slog"
	"net/http"
)

type UserService interface {
	Register(ctx context.Context)
	GetToken(ctx context.Context, username, pass string) (string, error)
}

type UserHandler struct {
	Service UserService
	Logger  *slog.Logger
}

func NewUserHandler(service UserService, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		Service: service,
		Logger:  logger,
	}
}

func (h *UserHandler) RegisterEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("POST /register", middleware.Middleware(h.Register))
	mux.HandleFunc("POST /register/", middleware.Middleware(h.Register))

	mux.HandleFunc("POST /get-token", middleware.Middleware(h.GetToken))
	mux.HandleFunc("POST /get-token/", middleware.Middleware(h.GetToken))
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {}
func (h *UserHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	token, err := h.Service.GetToken(r.Context(), "Jane", "")
	if err != nil {
		h.Logger.Error(err.Error())
	}

	w.Write([]byte(token))
}
