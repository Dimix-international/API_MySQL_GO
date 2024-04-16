package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	log *slog.Logger
}

func NewUserHandler(log *slog.Logger) *UserHandler {
	return &UserHandler{log: log}
}

func (h *UserHandler) RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

	// // admin routes
	// router.HandleFunc("/users/{userID}", auth.WithJWTAuth(h.handleGetUser, h.store)).Methods(http.MethodGet)
}

func (h *UserHandler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) handleRegister(w http.ResponseWriter, r *http.Request) {

}
