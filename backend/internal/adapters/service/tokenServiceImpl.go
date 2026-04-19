package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/config"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/ports"
)

type TokenServiceImpl struct {
	config    config.JWTConfig
	tokenRepo ports.RefreshTokenRepository
	userRepo  ports.UserRepository
}

func NewTokenService(cfg config.JWTConfig, tokenRepo ports.RefreshTokenRepository, userRepo ports.UserRepository) *TokenServiceImpl {
	return &TokenServiceImpl{
		config:    cfg,
		tokenRepo: tokenRepo,
		userRepo:  userRepo,
	}
}

func (t *TokenServiceImpl) GenerateTokenPair(ctx context.Context, userID uuid.UUID, email string) (*domain.TokenPair, error) {
	accessToken, expiresAt, err := t.generateAccessToken(userID, email)
	if err != nil {
		return nil, fmt.Errorf("generate access token: %w", err)
	}

	rawRefreshToken, err := t.generateAndStoreRefreshToken(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	return &domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: rawRefreshToken,
		ExpiresAt:    expiresAt.Unix(),
	}, nil
}

func (t *TokenServiceImpl) RefreshTokens(ctx context.Context, rawRefreshToken string) (*domain.TokenPair, error) {
	tokenHash := hashToken(rawRefreshToken)

	stored, err := t.tokenRepo.GetByTokenHash(ctx, tokenHash)
	if err != nil {
		return nil, domain.ErrTokenNotFound
	}

	if time.Now().After(stored.ExpiresAt) {
		_ = t.tokenRepo.DeleteByID(ctx, stored.ID)
		return nil, domain.ErrExpiredToken
	}

	if err := t.tokenRepo.DeleteByID(ctx, stored.ID); err != nil {
		return nil, fmt.Errorf("delete old refresh token: %w", err)
	}

	user, err := t.userRepo.GetByID(ctx, stored.UserID)
	if err != nil {
		return nil, fmt.Errorf("get user for refresh: %w", err)
	}

	return t.GenerateTokenPair(ctx, stored.UserID, user.Email)
}

func (t *TokenServiceImpl) ValidateAccessToken(tokenString string) (*domain.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.config.Secret), nil
	})

	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	claims, ok := token.Claims.(*domain.Claims)
	if !ok || !token.Valid {
		return nil, domain.ErrInvalidToken
	}

	return claims, nil
}

func (t *TokenServiceImpl) generateAccessToken(userID uuid.UUID, email string) (string, time.Time, error) {
	expiresAt := time.Now().Add(time.Duration(t.config.AccessTTL) * time.Minute)

	claims := domain.Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "todoapp",
			Subject:   userID.String(),
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

func (t *TokenServiceImpl) generateAndStoreRefreshToken(ctx context.Context, userID uuid.UUID) (string, error) {
	_ = t.tokenRepo.DeleteByUserID(ctx, userID)

	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("generate random bytes: %w", err)
	}

	rawToken := hex.EncodeToString(randomBytes)
	tokenHash := hashToken(rawToken)

	refreshToken := domain.RefreshToken{
		UserID:    userID,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(time.Duration(t.config.RefreshTTL) * 24 * time.Hour),
	}

	if err := t.tokenRepo.Create(ctx, refreshToken); err != nil {
		return "", fmt.Errorf("store refresh token: %w", err)
	}

	return rawToken, nil
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
