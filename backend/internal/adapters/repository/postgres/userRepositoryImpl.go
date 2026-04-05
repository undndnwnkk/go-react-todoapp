package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (u *UserRepositoryImpl) GetAll(ctx context.Context) ([]domain.User, error) {
	rows, err := u.db.QueryContext(ctx, `
		SELECT id, name, last_name, email, date_of_birth, password_hash, created_at
		FROM users
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]domain.User, 0)
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.LastName,
			&user.Email,
			&user.DateOfBirth,
			&user.PasswordHash,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var user domain.User

	err := u.db.QueryRowContext(
		ctx,
		`SELECT id, name, last_name, email, date_of_birth, password_hash, created_at
		 FROM users
		 WHERE id = $1`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Email,
		&user.DateOfBirth,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) Add(ctx context.Context, request domain.UserCreateRequest) (domain.UserIdResponse, error) {
	var response domain.UserIdResponse

	err := u.db.QueryRowContext(
		ctx,
		`INSERT INTO users
		 (name, last_name, email, date_of_birth, password_hash, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id`,
		request.Name,
		request.LastName,
		request.Email,
		request.DateOfBirth,
		request.PasswordHash,
		request.CreatedAt,
	).Scan(&response.ID)

	if err != nil {
		return domain.UserIdResponse{}, err
	}

	return response, nil
}

func (u *UserRepositoryImpl) UpdateById(ctx context.Context, id uuid.UUID, request domain.UserUpdateRequest) (domain.User, error) {
	var response domain.User

	err := u.db.QueryRowContext(
		ctx,
		`UPDATE users
		 SET name = $1, last_name = $2, email = $3, date_of_birth = $4, password_hash = $5
		 WHERE id = $6
		 RETURNING id, name, last_name, email, date_of_birth, password_hash, created_at`,
		request.Name,
		request.LastName,
		request.Email,
		request.DateOfBirth,
		request.PasswordHash,
		id,
	).Scan(
		&response.ID,
		&response.Name,
		&response.LastName,
		&response.Email,
		&response.DateOfBirth,
		&response.PasswordHash,
		&response.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return response, nil
}

func (u *UserRepositoryImpl) DeleteById(ctx context.Context, id uuid.UUID) error {
	res, err := u.db.ExecContext(
		ctx,
		`DELETE FROM users WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (u *UserRepositoryImpl) PatchById(ctx context.Context, id uuid.UUID, request domain.UserUpdateRequest) (domain.User, error) {
	var response domain.User

	err := u.db.QueryRowContext(
		ctx,
		`UPDATE users
		 SET name = $1, last_name = $2, email = $3, date_of_birth = $4, password_hash = $5
		 WHERE id = $6
		 RETURNING id, name, last_name, email, date_of_birth, password_hash, created_at`,
		request.Name,
		request.LastName,
		request.Email,
		request.DateOfBirth,
		request.PasswordHash,
		id,
	).Scan(
		&response.ID,
		&response.Name,
		&response.LastName,
		&response.Email,
		&response.DateOfBirth,
		&response.PasswordHash,
		&response.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return response, nil
}
