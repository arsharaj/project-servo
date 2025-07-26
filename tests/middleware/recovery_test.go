package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arsharaj/project-servo/middleware"
)

// TestRecoveryMiddleware verifies that the recovery middleware handles panics
// and responds with a 500 Internal Server Error.
func TestRecoveryMiddleware(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("simulated panic")
	})

	handler := middleware.Recovery(next)

	req := httptest.NewRequest("GET", "/panic", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected 500 status, got %d", w.Result().StatusCode)
	}
}
