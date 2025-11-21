package task

import (
	"os"
	"testing"
)

// TestMain controls the setup and teardown for our tests.
func TestMain(m *testing.M) {
	// Setup: Use a temporary file for testing
	fileName = "test_tasks.json"

	// Run tests
	code := m.Run()

	// Teardown: Clean up the test file
	os.Remove(fileName)

	os.Exit(code)
}

func TestAdd(t *testing.T) {
	// Clean up before test to ensure empty state
	os.Remove(fileName)

	title := "Test Task"
	task, err := Add(title)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if task.Title != title {
		t.Errorf("Expected title %q, got %q", title, task.Title)
	}

	if task.ID != 1 {
		t.Errorf("Expected ID 1, got %d", task.ID)
	}

	// Verify persistence
	tasks, err := List()
	if err != nil {
		t.Fatalf("Error listing tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

func TestComplete(t *testing.T) {
	// Setup
	os.Remove(fileName)
	Add("Task to complete")

	// Action
	err := Complete(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify
	tasks, _ := List()
	if !tasks[0].Done {
		t.Error("Expected task to be done, but it is not")
	}
}

func TestDelete(t *testing.T) {
	// Setup
	os.Remove(fileName)
	Add("Task to delete")

	// Action
	err := Delete(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify
	tasks, _ := List()
	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}
}
