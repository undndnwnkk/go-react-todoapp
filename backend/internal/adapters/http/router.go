package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/http/handler"
	authmw "github.com/undndnwnkk/go-react-todoapp/internal/adapters/http/middleware"
	"github.com/undndnwnkk/go-react-todoapp/internal/app"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/services"
)

func NewRouter(handlers *app.Handlers, tokenService services.TokenService) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	authMiddleware := authmw.AuthMiddleware(tokenService)

	router.Route("/api/v1", func(router chi.Router) {
		router.Get("/health", handler.Health)

		router.Route("/auth", func(router chi.Router) {
			router.Post("/register", handlers.AuthHandler.Register)
			router.Post("/login", handlers.AuthHandler.Login)
			router.Post("/refresh", handlers.AuthHandler.Refresh)
		})

		router.Route("/users", func(router chi.Router) {
			router.Get("/{id}", handlers.UserHandler.GetUser)
			router.Put("/{id}", handlers.UserHandler.UpdateUser)
			router.Delete("/{id}", handlers.UserHandler.DeleteUser)
		})

		router.Group(func(router chi.Router) {
			router.Use(authMiddleware)

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

			router.Route("/categories", func(router chi.Router) {
				router.Get("/", handlers.CategoryHandler.GetAllCategories)
				router.Post("/", handlers.CategoryHandler.CreateCategory)

				router.Route("/{id}", func(router chi.Router) {
					router.Put("/", handlers.CategoryHandler.UpdateCategory)
					router.Delete("/", handlers.CategoryHandler.DeleteCategory)
				})
			})
		})
	})

	return router
}
