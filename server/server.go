// Package server sets up and manages the HTTP server for the application.
package server

import (
	"context"
	"net/http"

	"github.com/arsharaj/project-servo/middleware"
)

// Server wraps the http.Server and provides methods to start and gracefully shut it down.
type Server struct {
	httpServer *http.Server
}

// New returns a new instance of Server with middleware and handlers configured.
func New(addr string) *Server {
	router := NewRouter()

	// Wrap with middleware chain : Recovery → Logging → Handlers
	handler := middleware.Recovery(middleware.Logging(router))

	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

// Start starts the HTTP server and blocks until it is stopped or an error occurs.
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// ShutDown gracefully shuts down the server using the provided context.
func (s *Server) ShutDown(ctx *context.Context) error {
	return s.httpServer.Shutdown(*ctx)
}

// ShutDownWithDefaultContext gracefully shuts down the server using a default background context.
func (s *Server) ShutDownWithDefaultContext() error {
	return s.httpServer.Shutdown(context.Background())
}
