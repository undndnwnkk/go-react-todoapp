package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/service/helpers"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/ports"
)

type CategoryServiceImpl struct {
	repo ports.CategoryRepository
}

func (c *CategoryServiceImpl) GetAll(ctx context.Context) ([]domain.Category, error) {
	categories, err := c.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *CategoryServiceImpl) Create(ctx context.Context, request domain.CategoryCreateRequest) (domain.Category, error) {
	if err := helpers.CheckCategoryCreateRequest(request); err != nil {
		return domain.Category{}, err
	}

	category, err := c.repo.Create(ctx, request)
	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (c *CategoryServiceImpl) GetByID(ctx context.Context, id uuid.UUID) (domain.Category, error) {
	category, err := c.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (c *CategoryServiceImpl) GetByUserID(ctx context.Context, userId uuid.UUID) ([]domain.Category, error) {
	category, err := c.repo.GetByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *CategoryServiceImpl) UpdateByID(ctx context.Context, id uuid.UUID, request domain.CategoryUpdateRequest) (domain.Category, error) {
	category, err := c.repo.UpdateByID(ctx, id, request)
	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (c *CategoryServiceImpl) PatchByID(ctx context.Context, id uuid.UUID, request domain.CategoryPatchRequest) (domain.Category, error) {
	category, err := c.repo.PatchById(ctx, id, request)
	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (c *CategoryServiceImpl) DeleteByID(ctx context.Context, id uuid.UUID) error {
	err := c.repo.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func NewCategoryService(repo ports.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{repo: repo}
}
