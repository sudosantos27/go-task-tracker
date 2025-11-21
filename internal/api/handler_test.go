package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sudosantos27/go-task-tracker/internal/store"
)

func TestGetTasks(t *testing.T) {
	// Setup
	repo := store.NewInMemoryTaskRepository()
	handler := NewHandler(repo)

	// Create a request
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler.GetTasks(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body (should be empty list initially)
	expected := "[]\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateTask(t *testing.T) {
	// Setup
	repo := store.NewInMemoryTaskRepository()
	handler := NewHandler(repo)

	// Create a request body
	taskData := map[string]string{
		"title":       "Test Task",
		"description": "Testing API",
	}
	body, _ := json.Marshal(taskData)

	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.CreateTask(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check response body contains the task
	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if response["title"] != "Test Task" {
		t.Errorf("handler returned unexpected title: got %v want %v",
			response["title"], "Test Task")
	}
}
