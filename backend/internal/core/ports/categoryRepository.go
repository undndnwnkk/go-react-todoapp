package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]domain.Category, error)
	Create(ctx context.Context, request domain.CategoryCreateRequest) (domain.Category, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Category, error)
	GetByUserID(ctx context.Context, userId uuid.UUID) ([]domain.Category, error)
	UpdateByID(ctx context.Context, id uuid.UUID, request domain.CategoryUpdateRequest) (domain.Category, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
	PatchByID(ctx context.Context, id uuid.UUID, request domain.CategoryPatchRequest) (domain.Category, error)
}
