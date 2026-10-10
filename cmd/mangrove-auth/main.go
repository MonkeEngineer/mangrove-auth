package main

import (
	"net/http"
	"os"
)

func main() {
	// Set up the HTTP handlers for health and readiness checks
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ready"}`))
	})

	// Get the port from the environment variable, default to 8080 if not set
	port := os.Getenv("MANGROVE_PORT")
	if port == "" {
		port = "8080"
	}

	// Start the HTTP server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
