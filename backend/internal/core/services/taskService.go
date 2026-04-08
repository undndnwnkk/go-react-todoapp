package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type TaskService interface {
	GetAll(ctx context.Context) ([]domain.Task, error)
	Create(ctx context.Context, request domain.TaskCreateRequest) (domain.Task, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Task, error)
	GetByUserID(ctx context.Context, userId uuid.UUID) ([]domain.Task, error)
	UpdateByID(ctx context.Context, id uuid.UUID, request domain.TaskUpdateRequest) (domain.Task, error)
	PatchByID(ctx context.Context, id uuid.UUID, request domain.TaskPatchRequest) (domain.Task, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
