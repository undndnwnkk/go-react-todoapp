package helpers

import (
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/http/handler"
)

type Handlers struct {
	UserHandler     *handler.UserHandler
	TaskHandler     *handler.TaskHandler
	CategoryHandler *handler.CategoryHandler
	AuthHandler     *handler.AuthHandler
}

func NewHandlers(u *handler.UserHandler, t *handler.TaskHandler, c *handler.CategoryHandler, a *handler.AuthHandler) *Handlers {
	return &Handlers{UserHandler: u, TaskHandler: t, CategoryHandler: c, AuthHandler: a}
}
