package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	Add(ctx context.Context, request domain.UserCreateRequest) (domain.UserIdResponse, error)
	UpdateById(ctx context.Context, id uuid.UUID, request domain.UserUpdateRequest) (domain.User, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
	PatchById(ctx context.Context, id uuid.UUID, request domain.UserUpdateRequest) (domain.User, error)
}
