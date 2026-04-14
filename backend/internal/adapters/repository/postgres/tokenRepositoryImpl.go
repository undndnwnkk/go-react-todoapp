package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type RefreshTokenRepositoryImpl struct {
	pool *pgxpool.Pool
}

func (r RefreshTokenRepositoryImpl) Create(ctx context.Context, request domain.RefreshToken) error {
	var id uuid.UUID
	err := r.pool.QueryRow(
		ctx,
		`INSERT INTO refresh_tokens (user_id, token_hash, expires_at, created_at)  
		VALUES ($1, $2, $3, $4) RETURNING id`,
		request.UserID,
		request.TokenHash,
		request.ExpiresAt,
		request.CreatedAt,
	).Scan(&id)

	if err != nil {
		return err
	}

	return nil
}

func (r RefreshTokenRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (domain.RefreshToken, error) {
	var token domain.RefreshToken
	err := r.pool.QueryRow(
		ctx,
		`SELECT * FROM refresh_tokens WHERE id = $1`,
		id,
	).Scan(&token.ID, &token.UserID, &token.TokenHash, &token.ExpiresAt, &token.CreatedAt)

	if err != nil {
		return domain.RefreshToken{}, err
	}

	return token, nil
}

func (r RefreshTokenRepositoryImpl) GetByUserID(ctx context.Context, userID uuid.UUID) (domain.RefreshToken, error) {
	var token domain.RefreshToken
	err := r.pool.QueryRow(
		ctx,
		`SELECT * FROM refresh_tokens WHERE user_id = $1`,
		userID,
	).Scan(&token.ID, &token.UserID, &token.TokenHash, &token.ExpiresAt, &token.CreatedAt)

	if err != nil {
		return domain.RefreshToken{}, err
	}

	return token, nil
}

func (r RefreshTokenRepositoryImpl) DeleteByID(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(
		ctx,
		`DELETE FROM refresh_tokens WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func NewRefreshTokenRepository(pool *pgxpool.Pool) *RefreshTokenRepositoryImpl {
	return &RefreshTokenRepositoryImpl{pool: pool}
}
