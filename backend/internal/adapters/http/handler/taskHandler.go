package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/services"
)

type TaskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(tsc services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: tsc}
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	tasks, err := h.taskService.GetAll(r.Context())
	if err != nil {
		http.Error(w, "error while get all request", http.StatusInternalServerError)
		return
	}

	err = encoder.Encode(tasks)
	if err != nil {
		http.Error(w, "invalid encoding", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	encoder := json.NewEncoder(w)

	var request domain.TaskCreateRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "invalid task create request", http.StatusBadRequest)
		return
	}

	task, err := h.taskService.Create(r.Context(), request)
	if err != nil {
		http.Error(w, "error while creating user", http.StatusInternalServerError)
		return
	}

	err = encoder.Encode(task)
	if err != nil {
		http.Error(w, "error while encoding task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	task, err := h.taskService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "error while get task by id", http.StatusInternalServerError)
		return
	}

	err = encoder.Encode(task)
	if err != nil {
		http.Error(w, "error while encoding task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var request domain.TaskUpdateRequest
	err = decoder.Decode(&request)
	if err != nil {
		http.Error(w, "error while decoding request", http.StatusBadRequest)
		return
	}

	task, err := h.taskService.UpdateByID(r.Context(), id, request)
	if err != nil {
		http.Error(w, "error while updating task", http.StatusBadRequest)
		return
	}

	err = encoder.Encode(task)
	if err != nil {
		http.Error(w, "error while encoding response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) PatchTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var request domain.TaskPatchRequest
	err = decoder.Decode(&request)
	if err != nil {
		http.Error(w, "error while decoding request", http.StatusBadRequest)
		return
	}

	task, err := h.taskService.PatchByID(r.Context(), id, request)
	if err != nil {
		http.Error(w, "error while patch task", http.StatusBadRequest)
		return
	}

	err = encoder.Encode(task)
	if err != nil {
		http.Error(w, "error while encoding response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.taskService.DeleteByID(r.Context(), id)
	if err != nil {
		http.Error(w, "error while deleting", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
