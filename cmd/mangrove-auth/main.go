package main

import (
	"log/slog"
	"net/http"
	"os"
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

	// Start the HTTP server
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		panic(err)
	}
}
