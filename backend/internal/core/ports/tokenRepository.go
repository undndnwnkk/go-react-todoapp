package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, request domain.RefreshToken) error
	GetByID(ctx context.Context, id uuid.UUID) (domain.RefreshToken, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (domain.RefreshToken, error)
	GetByTokenHash(ctx context.Context, tokenHash string) (domain.RefreshToken, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
}
