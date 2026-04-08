package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/service/helpers"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/ports"
)

type TaskServiceImpl struct {
	repo ports.TaskRepository
}

func (t *TaskServiceImpl) GetAll(ctx context.Context) ([]domain.Task, error) {
	tasks, err := t.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskServiceImpl) Create(ctx context.Context, request domain.TaskCreateRequest) (domain.Task, error) {
	if err := helpers.ValidateCreateRequest(request); err != nil {
		return domain.Task{}, err
	}

	task, err := t.repo.Create(ctx, request)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (t *TaskServiceImpl) GetByID(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	task, err := t.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (t *TaskServiceImpl) GetByUserID(ctx context.Context, userId uuid.UUID) ([]domain.Task, error) {
	task, err := t.repo.GetByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskServiceImpl) UpdateByID(ctx context.Context, id uuid.UUID, request domain.TaskUpdateRequest) (domain.Task, error) {
	task, err := t.repo.UpdateByID(ctx, id, request)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (t *TaskServiceImpl) PatchByID(ctx context.Context, id uuid.UUID, request domain.TaskPatchRequest) (domain.Task, error) {
	task, err := t.repo.PatchByID(ctx, id, request)
	if err != nil {
		return domain.Task{}, err
	}

	return task, err
}

func (t *TaskServiceImpl) DeleteByID(ctx context.Context, id uuid.UUID) error {
	err := t.repo.DeleteByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewTaskService(repo ports.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{repo: repo}
}
