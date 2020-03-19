package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/health", nil)
	handler := http.HandlerFunc(HealthCheck)
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("response returned wrong status: got %d want %d", response.Code, http.StatusOK)
	}

	if response.Body.String() != `{"alive": true}` {
		t.Errorf("response returned wrong body: got %q want %q", response.Body.String(), `{"alive": true}`)
	}
}
