package domain

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress        = "in_progress"
	StatusDone              = "done"
)

type Priority int

const (
	PriorityLow Priority = iota + 1
	PriorityMedium
	PriorityHigh
)

type Task struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	CategoryID  *uuid.UUID
	Title       string
	Description string
	Status      Status
	Priority    Priority
	DueDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskCreateRequest struct {
	UserID      uuid.UUID
	CategoryID  *uuid.UUID
	Title       string
	Description string
	Status      Status
	Priority    Priority
	DueDate     *time.Time
}

type TaskUpdateRequest struct {
	UserID      uuid.UUID
	CategoryID  *uuid.UUID
	Title       string
	Description string
	Status      Status
	Priority    Priority
	DueDate     *time.Time
}

type TaskPatchRequest struct {
	CategoryID  *uuid.UUID
	Title       *string
	Description *string
	Status      *Status
	Priority    *Priority
	DueDate     *time.Time
}
