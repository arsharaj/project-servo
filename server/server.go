package server

import (
	"context"
	"net/http"

	"github.com/arsharaj/project-servo/middleware"
)

type Server struct {
	httpServer *http.Server
}

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

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx *context.Context) error {
	return s.httpServer.Shutdown(*ctx)
}

func (s *Server) ShutDownWithDefaultContext() error {
	return s.httpServer.Shutdown(context.Background())
}
