package handler

import "github.com/undndnwnkk/go-react-todoapp/internal/core/services"

type CategoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(csc services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: csc}
}
