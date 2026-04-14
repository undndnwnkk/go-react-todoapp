package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/services"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(usc services.UserService) *UserHandler {
	return &UserHandler{userService: usc}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	decoder.DisallowUnknownFields()

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = encoder.Encode(user)
	if err != nil {
		http.Error(w, "Error while encoding response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	encoder := json.NewEncoder(w)

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var newUser domain.UserUpdateRequest
	err = decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "invalid user data", http.StatusBadRequest)
		return
	}

	user, err := h.userService.UpdateByID(r.Context(), id, newUser)
	if err != nil {
		http.Error(w, "invalid id or data", http.StatusBadRequest)
		return
	}

	if err = encoder.Encode(user); err != nil {
		http.Error(w, "Error while encoding response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.userService.DeleteByID(r.Context(), id)
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
