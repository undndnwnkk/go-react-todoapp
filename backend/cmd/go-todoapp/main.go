package main

import (
	"log"

	"github.com/undndnwnkk/go-react-todoapp/internal/app"
	"github.com/undndnwnkk/go-react-todoapp/internal/config"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error while loading config: %v", err)
	}

	application, err := app.New(cfg)
	if err != nil {
		log.Fatal("Error while creating app: ", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("app.Run: %v", err)
	}
}
