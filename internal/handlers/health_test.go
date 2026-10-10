package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/ecosystem-admin/mangrove-auth/internal/handlers"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	handlers.HealthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	if got := w.Body.String(); got != `{"status":"ok"}` {
		t.Errorf("unexpected body: %s", got)
	}
	if ct := w.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("unexpected Content-Type: %s", ct)
	}
}

func TestReadyHandler_Ready(t *testing.T) {
	var isShuttingDown atomic.Bool
	isShuttingDown.Store(false)

	req := httptest.NewRequest("GET", "/readyz", nil)
	w := httptest.NewRecorder()

	handlers.ReadyHandler(&isShuttingDown)(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	if got := w.Body.String(); got != `{"status":"ready"}` {
		t.Errorf("unexpected body: %s", got)
	}
}

func TestReadyHandler_ShuttingDown(t *testing.T) {
	var isShuttingDown atomic.Bool
	isShuttingDown.Store(true)

	req := httptest.NewRequest("GET", "/readyz", nil)
	w := httptest.NewRecorder()

	handlers.ReadyHandler(&isShuttingDown)(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("expected 503, got %d", w.Code)
	}
	if got := w.Body.String(); got != `{"status":"unavailable"}` {
		t.Errorf("unexpected body: %s", got)
	}
}
