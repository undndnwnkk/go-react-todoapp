package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type UserRepoImpl struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepoImpl {
	return &UserRepoImpl{
		pool: pool,
	}
}

func (u *UserRepoImpl) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	rows, err := u.pool.Query(ctx,
		`SELECT id, name, last_name, email, date_of_birth, password_hash, created_at FROM users`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.DateOfBirth, &user.PasswordHash, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepoImpl) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var user domain.User

	err := u.pool.QueryRow(
		ctx,
		`SELECT id, name, last_name, email, date_of_birth, password_hash, created_at FROM users WHERE id = $1`,
		id,
	).Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.DateOfBirth, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserRepoImpl) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	err := u.pool.QueryRow(
		ctx,
		`SELECT id, name, last_name, email, date_of_birth, password_hash, created_at FROM users WHERE email = $1`,
		email,
	).Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.DateOfBirth, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserRepoImpl) Create(ctx context.Context, request domain.UserCreateRequest) (domain.UserIdResponse, error) {
	var id uuid.UUID

	err := u.pool.QueryRow(
		ctx,
		`INSERT INTO users (name, last_name, email, date_of_birth, password_hash)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id`,
		request.Name,
		request.LastName,
		request.Email,
		request.DateOfBirth,
		request.Password,
	).Scan(&id)

	if err != nil {
		return domain.UserIdResponse{}, err
	}

	return domain.UserIdResponse{ID: id}, nil
}

func (u *UserRepoImpl) UpdateByID(ctx context.Context, id uuid.UUID, request domain.UserUpdateRequest) (domain.User, error) {
	var user domain.User

	err := u.pool.QueryRow(
		ctx,
		`UPDATE users
		SET name = $1, last_name = $2, email = $3, password_hash = $4, date_of_birth = $5
		WHERE id = $6
		RETURNING id, name, last_name, email, date_of_birth, password_hash, created_at`,
		request.Name,
		request.LastName,
		request.Email,
		request.PasswordHash,
		request.DateOfBirth,
		id,
	).Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.DateOfBirth, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserRepoImpl) DeleteByID(ctx context.Context, id uuid.UUID) error {
	res, err := u.pool.Exec(
		ctx,
		`DELETE FROM users WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (u *UserRepoImpl) PatchByID(ctx context.Context, id uuid.UUID, request domain.UserPatchRequest) (domain.User, error) {
	var base domain.User

	err := u.pool.QueryRow(
		ctx,
		`SELECT id, name, last_name, email, date_of_birth, password_hash, created_at FROM users WHERE id = $1`,
		id,
	).Scan(&base.ID, &base.Name, &base.LastName, &base.Email, &base.DateOfBirth, &base.PasswordHash, &base.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	toUpdate := validatePatchData(base, request)
	return u.UpdateByID(ctx, id, toUpdate)
}

func validatePatchData(base domain.User, request domain.UserPatchRequest) domain.UserUpdateRequest {
	var result domain.UserUpdateRequest

	if request.Name == nil {
		result.Name = base.Name
	} else {
		result.Name = *request.Name
	}
	if request.LastName == nil {
		result.LastName = base.LastName
	} else {
		result.LastName = *request.LastName
	}
	if request.Email == nil {
		result.Email = base.Email
	} else {
		result.Email = *request.Email
	}
	if request.PasswordHash == nil {
		result.PasswordHash = base.PasswordHash
	} else {
		result.PasswordHash = *request.PasswordHash
	}
	if request.DateOfBirth == nil && base.DateOfBirth != nil {
		result.DateOfBirth = base.DateOfBirth.Format("2006-01-02")
	} else if request.DateOfBirth != nil {
		result.DateOfBirth = *request.DateOfBirth
	}

	return result
}
