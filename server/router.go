package server

import (
	"net/http"

	"github.com/arsharaj/project-servo/handlers"
	"github.com/go-chi/chi/v5"
)

// NewRouter creates and returns a new HTTP router with all routes mounted.
func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Mount endpoints
	r.Get("/healthz", handlers.HealthCheck)

	return r
}
