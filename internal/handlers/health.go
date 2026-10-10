package handlers

import "net/http"

// HealthHandler handles the /healthz endpoint and returns a JSON response indicating the health status of the application.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
