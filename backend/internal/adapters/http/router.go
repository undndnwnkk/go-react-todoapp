package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/http/handler"
)

// TODO handlers to method args
func NewRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1", func(router chi.Router) {
		router.Get("/health", handler.Health)
	})

	return router
}
