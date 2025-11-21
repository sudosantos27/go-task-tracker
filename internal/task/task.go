package task

import "time"

// Task represents a task in our system.
type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// Repository defines the interface for task persistence.
type Repository interface {
	Create(t Task) (Task, error)
	GetAll() ([]Task, error)
}
