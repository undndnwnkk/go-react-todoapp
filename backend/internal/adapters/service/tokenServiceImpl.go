package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/config"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/ports"
)

type TokenServiceImpl struct {
	config          config.JWTConfig
	tokenRepository ports.RefreshTokenRepository
}

func (t *TokenServiceImpl) GenerateTokenPair(userID uuid.UUID, email string) (*domain.TokenPair, error) {
	accessToken, expiresAt, err := t.generateAccessToken(userID, email)
	if err != nil {
		return nil, errors.New(err)

	}
}

func NewTokenService(config config.JWTConfig, repo ports.RefreshTokenRepository) *TokenServiceImpl {
	return &TokenServiceImpl{config: config, tokenRepository: repo}
}
