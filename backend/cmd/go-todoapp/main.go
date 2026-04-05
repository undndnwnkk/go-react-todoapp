package main

import (
	"context"
	"log"
	http2 "net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/http"
	"github.com/undndnwnkk/go-react-todoapp/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error while loading config: %v", err)
	}

	// pgx logic
	ctx := context.Background()
	dsn := cfg.Database.GenerateDsn()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("Pgxpool.New:", err)
	}
	defer pool.Close()
	log.Println("Pgxpool.New pool created")

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Connection to database failed: ", err)
	}
	log.Print("Ping already working")

	router := http.NewRouter()
	err = http2.ListenAndServe(cfg.Server.Addr, router)
	if err != nil {
		log.Fatal("Error while serving: ", err)
	}
}
