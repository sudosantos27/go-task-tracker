package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

// Task represents a task in our system.
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

const fileName = "tasks.json"

func main() {
	// Check if any command was passed
	if len(os.Args) < 2 {
		fmt.Println("Usage: task <command> [arguments]")
		fmt.Println("Available commands: add, list, complete, delete")
		os.Exit(1)
	}

	// The first argument is the subcommand (add, list, etc.)
	switch os.Args[1] {
	case "add":
		// Define specific flags for the 'add' command
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("title", "", "Task title")

		// Parse arguments starting from the second one (os.Args[2:])
		addCmd.Parse(os.Args[2:])

		if *title == "" {
			fmt.Println("Error: Title is required. Usage: task add -title=\"My task\"")
			os.Exit(1)
		}

		handleAdd(*title)

	case "list":
		handleList()

	case "complete":
		completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
		id := completeCmd.Int("id", 0, "ID of the task to complete")

		completeCmd.Parse(os.Args[2:])

		if *id == 0 {
			fmt.Println("Error: ID is required. Usage: task complete -id=1")
			os.Exit(1)
		}

		handleComplete(*id)

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "ID of the task to delete")

		deleteCmd.Parse(os.Args[2:])

		if *id == 0 {
			fmt.Println("Error: ID is required. Usage: task delete -id=1")
			os.Exit(1)
		}

		handleDelete(*id)

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

// handleAdd creates a new task and saves it.
func handleAdd(title string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	// Generate a simple ID (last ID + 1)
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:        id,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)

	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Error saving task: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task added: %d: %s\n", newTask.ID, newTask.Title)
}

// handleList displays all tasks.
func handleList() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		fmt.Println("No pending tasks.")
		return
	}

	fmt.Println("Task list:")
	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		// Format: [ ] 1: Buy coffee (created: 2025-11-21 13:00)
		fmt.Printf("%s %d: %s (created: %s)\n", status, t.ID, t.Title, t.CreatedAt.Format("2006-01-02 15:04"))
	}
}

// handleComplete marks a task as completed.
func handleComplete(id int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Error: Task with ID %d not found\n", id)
		os.Exit(1)
	}

	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Error saving changes: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task %d marked as completed.\n", id)
}

// handleDelete deletes a task by its ID.
func handleDelete(id int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	index := -1
	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("Error: Task with ID %d not found\n", id)
		os.Exit(1)
	}

	// Remove element from slice: append(slice[:i], slice[i+1:]...)
	tasks = append(tasks[:index], tasks[index+1:]...)

	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Error saving changes: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task %d deleted.\n", id)
}

// loadTasks reads tasks from the JSON file.
func loadTasks() ([]Task, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // If it doesn't exist, return empty list
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
