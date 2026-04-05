package domain

import "github.com/google/uuid"

type Category struct {
	ID     uuid.UUID
	UserId uuid.UUID
	Name   string
	Color  string
}

type CategoryCreateRequest struct {
	UserId uuid.UUID
	Name   string
	Color  string
}

type CategoryUpdateRequest struct {
	UserId uuid.UUID
	Name   string
	Color  string
}
