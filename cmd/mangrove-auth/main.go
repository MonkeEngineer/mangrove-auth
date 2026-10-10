package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	// Listen for shutdown signals to gracefully shut down the server
	go func() {
		// Create a channel to listen for OS signals
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		// Wait for a signal to shut down
		<-quit

		// Log that the server is shutting down
		slog.Info("Shutting down server")

		// Create a context with a timeout for the shutdown process
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Attempt to shut down the server
		err := server.Shutdown(ctx)
		if err != nil {
			slog.Error("Server shutdown failed", "error", err)
		}
	}()

	// Start the HTTP server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start", "error", err)
	}

	// Log that the server has stopped
	slog.Info("Server stopped")
}
