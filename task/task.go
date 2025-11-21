package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Task represents a task in our system.
type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

var fileName = "tasks.json"

// Add creates a new task and saves it.
func Add(title, description string) (Task, error) {
	tasks, err := loadTasks()
	if err != nil {
		return Task{}, err
	}

	// Generate a simple ID (last ID + 1)
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:          id,
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	if err := saveTasks(tasks); err != nil {
		return Task{}, err
	}

	return newTask, nil
}

// List returns all tasks.
func List() ([]Task, error) {
	return loadTasks()
}

// Complete marks a task as completed.
func Complete(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			now := time.Now()
			tasks[i].CompletedAt = &now
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return saveTasks(tasks)
}

// Delete removes a task by its ID.
func Delete(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	index := -1
	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("task with ID %d not found", id)
	}

	// Remove element from slice
	tasks = append(tasks[:index], tasks[index+1:]...)

	return saveTasks(tasks)
}

// loadTasks reads tasks from the JSON file.
func loadTasks() ([]Task, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// saveTasks saves the list of tasks to the JSON file.
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
