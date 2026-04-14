package services

import (
	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type TokenService interface {
	GenerateTokenPair(userID uuid.UUID, email string) (*domain.TokenPair, error)
}
