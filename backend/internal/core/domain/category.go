package domain

import "github.com/google/uuid"

type Category struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Name   string
	Color  string
}

type CategoryCreateRequest struct {
	UserID uuid.UUID
	Name   string
	Color  string
}

type CategoryUpdateRequest struct {
	UserID uuid.UUID
	Name   string
	Color  string
}

type CategoryPatchRequest struct {
	UserID *uuid.UUID
	Name   *string
	Color  *string
}
