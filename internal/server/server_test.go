package server_test

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/ecosystem-admin/mangrove-auth/internal/server"
)

func TestSmokeHealthz(t *testing.T) {
	var isShuttingDown atomic.Bool
	srv := httptest.NewServer(server.NewRouter(&isShuttingDown))
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/healthz")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestSmokeReadyz(t *testing.T) {
	var isShuttingDown atomic.Bool
	srv := httptest.NewServer(server.NewRouter(&isShuttingDown))
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/readyz")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestSmokeReadyz_ShuttingDown(t *testing.T) {
	var isShuttingDown atomic.Bool
	srv := httptest.NewServer(server.NewRouter(&isShuttingDown))
	defer srv.Close()

	isShuttingDown.Store(true)

	resp, err := http.Get(srv.URL + "/readyz")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("expected 503, got %d", resp.StatusCode)
	}
}
