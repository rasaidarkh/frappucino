package handlers

import (
	"context"
	"fmt"
	"frappuccino/internal/handlers/middleware"
	"frappuccino/internal/helpers"
	"frappuccino/internal/models"
	"log/slog"
	"net/http"
)

type UserService interface {
	Register(ctx context.Context, user *models.User) (string, error)
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

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	pass := r.URL.Query().Get("password")

	token, err := h.Service.Register(
		r.Context(),
		&models.User{
			Username: username,
			Password: helpers.CreateMd5Hash(pass),
		},
	)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("error registering new user: %v", err))
		helpers.WriteError(w, http.StatusBadRequest, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Reponse{
		Key:   "successfully registered, here's your token",
		Value: token,
	})
}

func (h *UserHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	pass := r.URL.Query().Get("password")

	if len(username) == 0 || len(pass) == 0 {
		helpers.WriteError(w, http.StatusForbidden, fmt.Errorf("usesrname or password wasn't provided"))
		return
	}

	token, err := h.Service.GetToken(r.Context(), username, pass)
	if err != nil {
		h.Logger.Error(err.Error())
		helpers.WriteError(w, http.StatusForbidden, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Reponse{Key: "token", Value: token})
}
