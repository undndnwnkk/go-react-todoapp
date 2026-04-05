package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	// pgx logic
	ctx := context.Background()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("Pgxpool.New:", err)
	}
	defer pool.Close()
	log.Println("Pgxpool.New pool created")

	mux := http.NewServeMux()

	//mux.HandleFunc("/users", usersHandler)      // TODO
	//mux.HandleFunc("/users/{id}", usersHandler) // TODO

	http.ListenAndServe(":8080", mux)
}
