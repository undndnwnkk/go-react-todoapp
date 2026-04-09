package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type CategoryRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (r *CategoryRepositoryImpl) GetAll(ctx context.Context) ([]domain.Category, error) {
	const query = `
		SELECT id, user_id, name, color
		FROM categories
		ORDER BY name
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]domain.Category, 0)
	for rows.Next() {
		var category domain.Category

		if err := rows.Scan(
			&category.ID,
			&category.UserID,
			&category.Name,
			&category.Color,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepositoryImpl) Create(ctx context.Context, request domain.CategoryCreateRequest) (domain.Category, error) {
	const query = `
		INSERT INTO categories (user_id, name, color)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, name, color
	`

	var category domain.Category

	err := r.db.QueryRow(ctx, query,
		request.UserID,
		request.Name,
		request.Color,
	).Scan(
		&category.ID,
		&category.UserID,
		&category.Name,
		&category.Color,
	)
	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (r *CategoryRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (domain.Category, error) {
	const query = `
		SELECT id, user_id, name, color
		FROM categories
		WHERE id = $1
	`

	var category domain.Category

	err := r.db.QueryRow(ctx, query, id).Scan(
		&category.ID,
		&category.UserID,
		&category.Name,
		&category.Color,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Category{}, domain.ErrCategoryNotFound
		}
		return domain.Category{}, err
	}

	return category, nil
}

func (r *CategoryRepositoryImpl) GetByUserID(ctx context.Context, userId uuid.UUID) ([]domain.Category, error) {
	const query = `
		SELECT id, user_id, name, color
		FROM categories
		WHERE user_id = $1
`
	categories := make([]domain.Category, 0)

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category domain.Category
		if err := rows.Scan(
			&category.ID,
			&category.UserID,
			&category.Name,
			&category.Color,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepositoryImpl) UpdateByID(ctx context.Context, id uuid.UUID, request domain.CategoryUpdateRequest) (domain.Category, error) {
	const query = `
		UPDATE categories
		SET user_id = $2,
			name = $3,
			color = $4
		WHERE id = $1
		RETURNING id, user_id, name, color
	`

	var category domain.Category

	err := r.db.QueryRow(ctx, query,
		id,
		request.UserID,
		request.Name,
		request.Color,
	).Scan(
		&category.ID,
		&category.UserID,
		&category.Name,
		&category.Color,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Category{}, domain.ErrCategoryNotFound
		}
		return domain.Category{}, err
	}

	return category, nil
}

func (r *CategoryRepositoryImpl) DeleteByID(ctx context.Context, id uuid.UUID) error {
	const query = `
		DELETE FROM categories
		WHERE id = $1
	`

	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return domain.ErrCategoryNotFound
	}

	return nil
}

func (r *CategoryRepositoryImpl) PatchByID(ctx context.Context, id uuid.UUID, request domain.CategoryPatchRequest) (domain.Category, error) {
	category, err := r.GetByID(ctx, id)
	if err != nil {
		return domain.Category{}, err
	}

	toUpdate := r.validatePatchData(category, request)
	updated, err := r.UpdateByID(ctx, id, toUpdate)
	if err != nil {
		return domain.Category{}, err
	}

	return updated, nil
}

func (r *CategoryRepositoryImpl) validatePatchData(base domain.Category, request domain.CategoryPatchRequest) domain.CategoryUpdateRequest {
	var result domain.CategoryUpdateRequest

	if request.UserID == nil {
		result.UserID = base.UserID
	} else {
		result.UserID = *request.UserID
	}
	if request.Name == nil {
		result.Name = base.Name
	} else {
		result.Name = *request.Name
	}
	if request.Color == nil {
		result.Color = base.Color
	} else {
		result.Color = *request.Color
	}
	return result
}
