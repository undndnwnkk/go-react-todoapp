package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	http2 "github.com/undndnwnkk/go-react-todoapp/internal/adapters/http"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/http/handler"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/repository/postgres"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/service"
	"github.com/undndnwnkk/go-react-todoapp/internal/config"
	"github.com/undndnwnkk/go-react-todoapp/internal/helpers"
)

type App struct {
	cfg    *config.Config
	server *http.Server
	db     *pgxpool.Pool
}

func New(cfg *config.Config) (*App, error) {
	ctx := context.Background()
	dsn := cfg.Database.GenerateDsn()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		//log.Fatal("Pgxpool.New: ", err)
		return nil, fmt.Errorf("Pgxpool.New: %w", err)
	}
	slog.Info("Pgxpool.New pool created")

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Connection to database failed: %w", err)
	}
	slog.Info("Ping already working")

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

	handlers := helpers.NewHandlers(userHandler, taskHandler, categoryHandler, authHandler)

	router := http2.NewRouter(handlers, tokenService)

	server := &http.Server{
		Addr:              cfg.Server.Addr,
		Handler:           router,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return &App{
		cfg:    cfg,
		server: server,
		db:     pool,
	}, nil
}

func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	serverErr := make(chan error, 1)
	go func() {
		slog.Info("http server starting", "addr", a.cfg.Server.Addr)
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- fmt.Errorf("http server: %w", err)
		}
	}()

	select {
	case err := <-serverErr:
		return err
	case <-ctx.Done():
		slog.Info("shutdown initialized")
	}

	return a.shutdown()
}

func (a *App) shutdown() error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*30,
	)
	defer cancel()

	var shutdownErr error

	slog.Info("stopping http server...")
	if err := a.server.Shutdown(ctx); err != nil {
		shutdownErr = fmt.Errorf("http shutdown: %w", err)
		slog.Error("http server shutdown error", "error", err)
	} else {
		slog.Info("http server stopped")
	}

	a.db.Close()
	return shutdownErr
}
