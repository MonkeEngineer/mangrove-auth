package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/ecosystem-admin/mangrove-auth/internal/server"
)

func main() {
	// Initialize an atomic boolean to track if the server is shutting down
	var isShuttingDown atomic.Bool
	isShuttingDown.Store(false)

	// Get the port from the environment variable, default to 8080 if not set
	port := os.Getenv("MANGROVE_PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server and log the port it's running on
	slog.Info("Starting server", "port", port)
	server := server.NewServer(&isShuttingDown, port)

	// Listen for shutdown signals to gracefully shut down the server
	go func() {
		// Create a channel to listen for OS signals
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		// Wait for a signal to shut down
		<-quit
		isShuttingDown.Store(true)

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
