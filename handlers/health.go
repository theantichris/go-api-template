package handlers

import (
	"io"
	"net/http"
)

// HealthCheck will return an "alive" response if the server is up.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}
