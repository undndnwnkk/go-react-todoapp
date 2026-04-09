package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/http/handler"
	"github.com/undndnwnkk/go-react-todoapp/internal/app"
)

// TODO handlers to method args
func NewRouter(handlers *app.Handlers) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1", func(router chi.Router) {
		router.Get("/health", handler.Health)

		// Auth TODO
		router.Route("/auth", func(router chi.Router) {
			router.Post("/register", handlers.AuthHandler.Pass)
			router.Post("/login", handlers.AuthHandler.Pass)
			router.Post("/refresh", handlers.AuthHandler.Pass)
		})

		// Users
		router.Route("/users", func(router chi.Router) {
			router.Get("/{id}", handlers.UserHandler.GetUser)
			router.Put("/{id}", handlers.UserHandler.UpdateUser)
			router.Delete("/{id}", handlers.UserHandler.DeleteUser)
		})

		// Tasks
		router.Route("/tasks", func(router chi.Router) {
			router.Get("/", handlers.TaskHandler.GetAll)
			router.Post("/", handlers.TaskHandler.CreateTask)

			router.Route("/{id}", func(router chi.Router) {
				router.Get("/", handlers.TaskHandler.GetTask)
				router.Put("/", handlers.TaskHandler.UpdateTask)
				router.Patch("/", handlers.TaskHandler.PatchTask)
				router.Delete("/", handlers.TaskHandler.DeleteTask)
			})
		})

		// Categories
		router.Route("/categories", func(router chi.Router) {
			router.Get("/", handlers.CategoryHandler.GetAllCategories)
			router.Post("/", handlers.CategoryHandler.CreateCategory)

			router.Route("/{id}", func(router chi.Router) {
				router.Put("/", handlers.CategoryHandler.UpdateCategory)
				router.Delete("/", handlers.CategoryHandler.DeleteCategory)
			})
		})
	})

	return router
}
