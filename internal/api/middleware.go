package api

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the method, path, and duration of each request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log after the request is handled
		duration := time.Since(start)
		log.Printf("[%s] %s - %v", r.Method, r.URL.Path, duration)
	})
}
