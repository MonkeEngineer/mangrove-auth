package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create a new HTTP server mux
	mux := http.NewServeMux()

	// Set up the HTTP handlers for health and readiness checks
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ready"}`))
	})

	// Get the port from the environment variable, default to 8080 if not set
	port := os.Getenv("MANGROVE_PORT")
	if port == "" {
		port = "8080"
	}

	// Log the port the server is listening on
	slog.Info("Starting server", "port", port)

	// Create the HTTP server instance
	server := &http.Server{
		Addr:         "0.0.0.0:" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start the HTTP server
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
