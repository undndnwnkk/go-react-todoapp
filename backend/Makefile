-include .env
export

MIGRATE_DIR := ./migrations
DB_DSN ?= postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
BINARY_NAME := go-todoapp
BUILD_DIR := bin
LINT_FLAGS := --timeout 5m

ifndef NAME
NAME :=
endif

run:
	@echo "==> Starting go application"
	@go run ./cmd/go-todoapp

build:
	@echo "==> Building $(BINARY_NAME)"
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/go-todoapp ./cmd/go-todoapp

clean:
	@echo "==> Cleaning"
	@rm -rf $(BUILD_DIR)
	@go clean -cache -testcache
	@echo "==> Clean complete"

test:
	@echo "==> Testing application"
	@go test ./...

lint:
	@echo "==> Running linter"
	@which golangci-lint > /dev/null || (echo "golangci-lint not found. Run: make install-tools" && exit 1)
	golangci-lint run $(LINT_FLAGS) ./...

migrate-up:
	@echo "==> Running migrations up"
	@which goose > /dev/null || (echo "goose not found"; exit 1)
	goose -dir $(MIGRATE_DIR) postgres "$(DB_DSN)" up

migrate-down:
	@echo "==> Rolling back last migration"
	@which goose > /dev/null || (echo "goose not found"; exit 1)
	goose -dir $(MIGRATE_DIR) postgres "$(DB_DSN)" down

migrate-status:
	@echo "==> Migration status"
	@which goose > /dev/null || (echo "goose not found"; exit 1)
	goose -dir $(MIGRATE_DIR) postgres "$(DB_DSN)" status

migrate-create:
	@if [ -z "$(NAME)" ]; then echo "NAME is required. Usage: make migrate-create NAME=migration_name"; exit 1; fi
	@echo "==> Creating migration $(NAME)"
	goose -dir $(MIGRATE_DIR) create $(NAME) sql

migrate-reset:
	@echo "==> Resetting all migrations (this will DROP all tables!)"
	@read -p "Are you sure? [y/N] " confirm && [ "$$confirm" = "y" ]
	goose -dir $(MIGRATE_DIR) postgres "$(DB_DSN)" reset

compose-up:
	@echo "==> Running docker compose"
	@docker compose up -d

compose-down:
	@echo "==> Running docker compose down"
	@docker compose down

install-tools:
	@echo "==> Installing development tools"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/air-verse/air@latest
	go install github.com/vektra/mockery/v2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "==> Tools installed"