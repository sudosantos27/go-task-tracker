package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sudosantos27/go-task-tracker/internal/task"
)

// Handler holds dependencies for HTTP handlers.
type Handler struct {
	repo task.Repository
}

// NewHandler creates a new Handler instance.
func NewHandler(repo task.Repository) *Handler {
	return &Handler{repo: repo}
}

// GetTasks handles GET /tasks.
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// CreateTask handles POST /tasks.
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if input.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	newTask := task.Task{
		Title:       input.Title,
		Description: input.Description,
		Done:        false,
		CreatedAt:   time.Now(),
	}

	created, err := h.repo.Create(newTask)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
