package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
		return nil, err
	}

	refreshToken, err := t.generateRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	return &domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt.Unix(),
	}, nil
}

func (t *TokenServiceImpl) generateAccessToken(userID uuid.UUID, email string) (string, time.Time, error) {

	expiresAt := time.Now().Add(time.Duration(t.config.AccessTTL) * time.Minute)

	claims := domain.Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "todoapp",
			Subject:   fmt.Sprintf("%d", userID),
			Audience:  jwt.ClaimStrings{"todoapp-api"},
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(t.config.Secret))
	if err != nil {
		return "", time.Time{}, err
	}

	return signed, expiresAt, nil
}

func (t *TokenServiceImpl) generateRefreshToken(userID uuid.UUID) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(t.config.RefreshTTL) * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        uuid.New().String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.config.Secret)
}

func NewTokenService(config config.JWTConfig, repo ports.RefreshTokenRepository) *TokenServiceImpl {
	return &TokenServiceImpl{config: config, tokenRepository: repo}
}
