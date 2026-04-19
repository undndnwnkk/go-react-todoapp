package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type RefreshTokenRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewRefreshTokenRepository(pool *pgxpool.Pool) *RefreshTokenRepositoryImpl {
	return &RefreshTokenRepositoryImpl{pool: pool}
}

func (r *RefreshTokenRepositoryImpl) Create(ctx context.Context, request domain.RefreshToken) error {
	_, err := r.pool.Exec(
		ctx,
		`INSERT INTO refresh_tokens (user_id, token_hash, expires_at)
		 VALUES ($1, $2, $3)`,
		request.UserID,
		request.TokenHash,
		request.ExpiresAt,
	)
	return err
}

func (r *RefreshTokenRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (domain.RefreshToken, error) {
	var token domain.RefreshToken
	err := r.pool.QueryRow(
		ctx,
		`SELECT id, user_id, token_hash, expires_at, created_at
		 FROM refresh_tokens WHERE id = $1`,
		id,
	).Scan(&token.ID, &token.UserID, &token.TokenHash, &token.ExpiresAt, &token.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.RefreshToken{}, domain.ErrTokenNotFound
		}
		return domain.RefreshToken{}, err
	}

	return token, nil
}

func (r *RefreshTokenRepositoryImpl) GetByUserID(ctx context.Context, userID uuid.UUID) (domain.RefreshToken, error) {
	var token domain.RefreshToken
	err := r.pool.QueryRow(
		ctx,
		`SELECT id, user_id, token_hash, expires_at, created_at
		 FROM refresh_tokens WHERE user_id = $1`,
		userID,
	).Scan(&token.ID, &token.UserID, &token.TokenHash, &token.ExpiresAt, &token.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.RefreshToken{}, domain.ErrTokenNotFound
		}
		return domain.RefreshToken{}, err
	}

	return token, nil
}

func (r *RefreshTokenRepositoryImpl) GetByTokenHash(ctx context.Context, tokenHash string) (domain.RefreshToken, error) {
	var token domain.RefreshToken
	err := r.pool.QueryRow(
		ctx,
		`SELECT id, user_id, token_hash, expires_at, created_at
		 FROM refresh_tokens WHERE token_hash = $1`,
		tokenHash,
	).Scan(&token.ID, &token.UserID, &token.TokenHash, &token.ExpiresAt, &token.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.RefreshToken{}, domain.ErrTokenNotFound
		}
		return domain.RefreshToken{}, err
	}

	return token, nil
}

func (r *RefreshTokenRepositoryImpl) DeleteByID(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(
		ctx,
		`DELETE FROM refresh_tokens WHERE id = $1`,
		id,
	)
	return err
}

func (r *RefreshTokenRepositoryImpl) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	_, err := r.pool.Exec(
		ctx,
		`DELETE FROM refresh_tokens WHERE user_id = $1`,
		userID,
	)
	return err
}
