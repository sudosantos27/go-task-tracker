package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sudosantos27/go-task-tracker/internal/api"
	"github.com/sudosantos27/go-task-tracker/internal/store"
)

func main() {
	// Initialize dependencies
	repo := store.NewInMemoryTaskRepository()
	handler := api.NewHandler(repo)

	// Define handlers
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetTasks(w, r)
		case http.MethodPost:
			handler.CreateTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

// healthHandler returns a simple JSON status.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write status code
	w.WriteHeader(http.StatusOK)

	// Write JSON response
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
