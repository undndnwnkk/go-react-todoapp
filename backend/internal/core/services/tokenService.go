package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type TokenService interface {
	GenerateTokenPair(ctx context.Context, userID uuid.UUID, email string) (*domain.TokenPair, error)
	RefreshTokens(ctx context.Context, rawRefreshToken string) (*domain.TokenPair, error)
	ValidateAccessToken(tokenString string) (*domain.Claims, error)
}
