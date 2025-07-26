// Package handlers_test contains tests for HTTP handler functions.
package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arsharaj/project-servo/handlers"
)

// TestHealthCheck verifies that the health check endpoint returns 200 OK
// and the response is in application/json format.
func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()

	handlers.HealthCheck(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}
}
