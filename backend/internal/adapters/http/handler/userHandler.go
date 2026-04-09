package handler

import "github.com/undndnwnkk/go-react-todoapp/internal/core/services"

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(usc services.UserService) *UserHandler {
	return &UserHandler{userService: usc}
}
