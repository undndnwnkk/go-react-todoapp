package main

import (
	"context"
	"log"
	http2 "net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/http"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/http/handler"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/repository/postgres"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/service"
	"github.com/undndnwnkk/go-react-todoapp/internal/app"
	"github.com/undndnwnkk/go-react-todoapp/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error while loading config: %v", err)
	}

	ctx := context.Background()
	dsn := cfg.Database.GenerateDsn()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("Pgxpool.New: ", err)
	}
	defer pool.Close()
	log.Println("Pgxpool.New pool created")

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Connection to database failed: ", err)
	}
	log.Println("Ping already working")

	userRepository := postgres.NewUserRepository(pool)
	taskRepository := postgres.NewTaskRepository(pool)
	categoryRepository := postgres.NewCategoryRepository(pool)
	tokenRepository := postgres.NewRefreshTokenRepository(pool)

	userService := service.NewUserService(userRepository)
	taskService := service.NewTaskService(taskRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	tokenService := service.NewTokenService(cfg.JWT, tokenRepository, userRepository)

	userHandler := handler.NewUserHandler(userService)
	taskHandler := handler.NewTaskHandler(taskService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	authHandler := handler.NewAuthHandler(userService, tokenService)

	handlers := app.NewHandlers(userHandler, taskHandler, categoryHandler, authHandler)

	router := http.NewRouter(handlers, tokenService)

	log.Printf("Starting server on %s", cfg.Server.Addr)
	if err := http2.ListenAndServe(cfg.Server.Addr, router); err != nil {
		log.Fatal("Error while serving: ", err)
	}
}
