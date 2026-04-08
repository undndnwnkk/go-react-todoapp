package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type UserService interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	Register(ctx context.Context, request domain.UserCreateRequest) (domain.UserIdResponse, error)
	Login(ctx context.Context, request domain.UserLoginRequest) (domain.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	UpdateByID(ctx context.Context, id uuid.UUID, request domain.UserUpdateRequest) (domain.User, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
