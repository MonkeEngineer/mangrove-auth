package server

import (
	"net/http"
	"sync/atomic"
	"time"
)

func NewServer(isShuttingDown *atomic.Bool, port string) *http.Server {
	router := NewRouter(isShuttingDown)

	return &http.Server{
		Addr:         "0.0.0.0:" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}
