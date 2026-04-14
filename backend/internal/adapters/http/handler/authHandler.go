package handler

import (
	"encoding/json"
	"net/http"

	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/services"
)

type AuthHandler struct {
	userService services.UserService
}

func NewAuthHandler(authService string) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// TODO delete this shit
func (a *AuthHandler) Pass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("hi")
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request domain.UserCreateRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "invalid create request", http.StatusBadRequest)
		return
	}

	id, err := h.userService.Register(r.Context(), request)
	if err != nil {
		http.Error(w, "invalid create request", http.StatusBadRequest)
		return
	}

}
