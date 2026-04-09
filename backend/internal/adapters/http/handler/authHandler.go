package handler

import (
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	// TODO
	authService string
}

func NewAuthHandler(authService string) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (a *AuthHandler) Pass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("hi")
	w.WriteHeader(http.StatusOK)
}
