package handler

import "github.com/undndnwnkk/go-react-todoapp/internal/core/services"

type TaskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(tsc services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: tsc}
}
