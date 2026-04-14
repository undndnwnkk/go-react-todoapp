package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/services"
)

type CategoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(csc services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: csc}
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	categories, err := h.categoryService.GetAll(r.Context())
	if err != nil {
		http.Error(w, "error while get all", http.StatusInternalServerError)
		return
	}

	err = decoder.Decode(&categories)
	if err != nil {
		http.Error(w, "error while decoding", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request domain.CategoryCreateRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "error while decoding", http.StatusBadRequest)
		return
	}

	category, err := h.categoryService.Create(r.Context(), request)
	if err != nil {
		http.Error(w, "error while creating category", http.StatusBadRequest)
		return
	}

	err = encoder.Encode(category)
	if err != nil {
		http.Error(w, "error while encoding", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var request domain.CategoryUpdateRequest
	err = decoder.Decode(&request)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	category, err := h.categoryService.UpdateByID(r.Context(), id, request)
	if err != nil {
		http.Error(w, "category with this id not found", http.StatusBadRequest)
		return
	}

	err = encoder.Encode(category)
	if err != nil {
		http.Error(w, "error while encoding", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.categoryService.DeleteByID(r.Context(), id)
	if err != nil {
		http.Error(w, "id not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
