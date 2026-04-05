-include .env
export

MIGRATE_DIR := ./migrations
DB_DSN ?= postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

ifndef NAME
NAME :=
endif

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