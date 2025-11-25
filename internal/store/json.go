package store

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/sudosantos27/go-task-tracker/internal/task"
)

// JSONTaskRepository implements task.Repository using a JSON file.
type JSONTaskRepository struct {
	mu       sync.Mutex
	filename string
}

// NewJSONTaskRepository creates a new instance.
func NewJSONTaskRepository(filename string) *JSONTaskRepository {
	return &JSONTaskRepository{
		filename: filename,
	}
}

// Create adds a new task to the repository.
func (r *JSONTaskRepository) Create(t task.Task) (task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.load()
	if err != nil {
		return task.Task{}, err
	}

	// Find max ID to simulate auto-increment
	maxID := 0
	for _, existing := range tasks {
		if existing.ID > maxID {
			maxID = existing.ID
		}
	}
	t.ID = maxID + 1

	tasks = append(tasks, t)

	if err := r.save(tasks); err != nil {
		return task.Task{}, err
	}

	return t, nil
}

// GetAll returns all tasks.
func (r *JSONTaskRepository) GetAll() ([]task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.load()
}

// load reads tasks from the file.
func (r *JSONTaskRepository) load() ([]task.Task, error) {
	file, err := os.ReadFile(r.filename)
	if os.IsNotExist(err) {
		return []task.Task{}, nil
	}
	if err != nil {
		return nil, err
	}

	if len(file) == 0 {
		return []task.Task{}, nil
	}

	var tasks []task.Task
	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// save writes tasks to the file.
func (r *JSONTaskRepository) save(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filename, data, 0644)
}
