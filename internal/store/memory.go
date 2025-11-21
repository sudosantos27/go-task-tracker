package store

import (
	"sync"

	"github.com/sudosantos27/go-task-tracker/internal/task"
)

// InMemoryTaskRepository implements task.Repository using a slice.
type InMemoryTaskRepository struct {
	mu     sync.RWMutex
	tasks  []task.Task
	nextID int
}

// NewInMemoryTaskRepository creates a new instance.
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  []task.Task{},
		nextID: 1,
	}
}

// Create adds a new task to the repository.
func (r *InMemoryTaskRepository) Create(t task.Task) (task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t.ID = r.nextID
	r.nextID++
	r.tasks = append(r.tasks, t)
	return t, nil
}

// GetAll returns all tasks.
func (r *InMemoryTaskRepository) GetAll() ([]task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Return a copy to avoid race conditions if the caller modifies the result
	results := make([]task.Task, len(r.tasks))
	copy(results, r.tasks)
	return results, nil
}
