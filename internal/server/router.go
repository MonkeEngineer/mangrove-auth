package server

import (
	"net/http"
	"sync/atomic"

	"github.com/ecosystem-admin/mangrove-auth/internal/handlers"
)

func NewRouter(isShuttingDown *atomic.Bool) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/healthz", handlers.HealthHandler)
	router.HandleFunc("/readyz", handlers.ReadyHandler(isShuttingDown))

	return router
}
