package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	Create(ctx context.Context, request domain.UserCreateRequest) (domain.UserIdResponse, error)
	UpdateByID(ctx context.Context, id uuid.UUID, request domain.UserUpdateRequest) (domain.User, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
	PatchByID(ctx context.Context, id uuid.UUID, request domain.UserPatchRequest) (domain.User, error)
}
