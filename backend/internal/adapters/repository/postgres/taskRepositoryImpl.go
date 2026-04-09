package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type TaskRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewTaskRepositoryImpl(pool *pgxpool.Pool) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{pool: pool}
}

func (t *TaskRepositoryImpl) GetAll(ctx context.Context) ([]domain.Task, error) {
	result := make([]domain.Task, 0)
	rows, err := t.pool.Query(
		ctx,
		`SELECT * FROM tasks`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.CategoryID, &task.Title, &task.Description, &task.Status, &task.Priority, &task.DueDate, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, task)
	}
	return result, nil
}

func (t *TaskRepositoryImpl) Create(ctx context.Context, request domain.TaskCreateRequest) (domain.Task, error) {
	var res domain.Task
	err := t.pool.QueryRow(
		ctx,
		`INSERT INTO tasks (user_id, category_id, title, description, status, priority, due_date)
    	VALUES ($1, $2, $3, $4, $5, $6, $7)
    	RETURNING id, user_id, category_id, title, description, status, priority, due_date, created_at, updated_at`,
		request.UserID,
		request.CategoryID,
		request.Title,
		request.Description,
		request.Status,
		request.Priority,
		request.DueDate,
	).Scan(&res.ID, &res.UserID, &res.CategoryID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return domain.Task{}, err
	}

	return res, nil
}

func (t *TaskRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	var res domain.Task
	err := t.pool.QueryRow(
		ctx,
		`SELECT * FROM tasks WHERE id = $1`,
		id,
	).Scan(&res.ID, &res.UserID, &res.CategoryID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return domain.Task{}, err
	}

	return res, nil
}

func (t *TaskRepositoryImpl) GetByUserID(ctx context.Context, user_id uuid.UUID) ([]domain.Task, error) {
	result := make([]domain.Task, 0)
	rows, err := t.pool.Query(
		ctx,
		`SELECT * FROM tasks WHERE user_id = $1`,
		user_id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var res domain.Task
		if err := rows.Scan(&res.ID, &res.UserID, &res.CategoryID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.DueDate, &res.CreatedAt, &res.UpdatedAt); err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return result, nil
}

func (t *TaskRepositoryImpl) UpdateByID(ctx context.Context, id uuid.UUID, request domain.TaskUpdateRequest) (domain.Task, error) {
	var res domain.Task
	err := t.pool.QueryRow(
		ctx,
		`UPDATE tasks 
		SET category_id = $1, title = $2, description = $3, status = $4, priority = $5, due_date = $6 
		WHERE id = $7 
		RETURNING id, user_id, category_id, title, description, status, priority, due_date, created_at, updated_at`,
		request.CategoryID,
		request.Title,
		request.Description,
		request.Status,
		request.Priority,
		request.DueDate,
		id,
	).Scan(&res.ID, &res.UserID, &res.CategoryID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return domain.Task{}, err
	}

	return res, nil
}

func (t *TaskRepositoryImpl) DeleteByID(ctx context.Context, id uuid.UUID) error {
	row, err := t.pool.Exec(
		ctx,
		`DELETE FROM tasks WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	affected := row.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return domain.ErrTaskNotFound
	}
	return nil
}

func (t *TaskRepositoryImpl) PatchByID(ctx context.Context, id uuid.UUID, request domain.TaskPatchRequest) (domain.Task, error) {
	var res domain.Task
	err := t.pool.QueryRow(
		ctx,
		`SELECT * FROM tasks WHERE id = $1`,
		id,
	).Scan(&res.ID, &res.UserID, &res.CategoryID, &res.Title, &res.Description, &res.Status, &res.Priority, &res.DueDate, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return domain.Task{}, err
	}

	toUpdate := validateTaskPatchData(res, request)
	return t.UpdateByID(ctx, id, toUpdate)
}

func validateTaskPatchData(base domain.Task, request domain.TaskPatchRequest) domain.TaskUpdateRequest {
	var result domain.TaskUpdateRequest
	if request.CategoryID == nil {
		result.CategoryID = base.CategoryID
	} else {
		result.CategoryID = request.CategoryID
	}
	if request.Title == nil {
		result.Title = base.Title
	} else {
		result.Title = *request.Title
	}
	if request.Description == nil {
		result.Description = base.Description
	} else {
		result.Description = *request.Description
	}
	if request.Status == nil {
		result.Status = base.Status
	} else {
		result.Status = *request.Status
	}
	if request.Priority == nil {
		result.Priority = base.Priority
	} else {
		result.Priority = *request.Priority
	}
	if request.DueDate == nil {
		result.DueDate = base.DueDate
	}

	return result
}

func NewTaskRepository(pool *pgxpool.Pool) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{pool: pool}
}
