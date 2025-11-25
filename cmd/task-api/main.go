package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sudosantos27/go-task-tracker/internal/api"
	"github.com/sudosantos27/go-task-tracker/internal/store"
)

func main() {
	// 1. Configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 2. Dependencies
	repo := store.NewJSONTaskRepository("tasks.json")
	handler := api.NewHandler(repo)

	// 3. Router & Middleware
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetTasks(w, r)
		case http.MethodPost:
			handler.CreateTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 4. Server Setup
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: api.LoggingMiddleware(mux),
	}

	// 5. Start Server in Goroutine
	go func() {
		log.Printf("Starting server on :%s...\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %s\n", err)
		}
	}()

	// 6. Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // Wait for signal

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	log.Println("Server exited properly")
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
