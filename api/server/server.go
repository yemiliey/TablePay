package server

import (
	"context"
	"net"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/pkg/errors"
)

// Server is a Table Pay API server.
type Server struct {
	service    *goa.Service
	listener   net.Listener
	httpServer *http.Server
}

// New creates a new Server.
func New(listener net.Listener, posDataAPIClient servicepb.DataAPIClient, logger hatchet.Logger) *Server {

	// Create service
	service := goa.New("Table Pay API")
	service.WithLogger(goalog.Adapt(logger))

	// Mount middleware
	service.Use(statusMiddleware())
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "assets" controller
	c := NewAssetsController(service)
	app.MountAssetsController(service, c)
	// Mount "tablepay" controller
	c2 := NewTablepayController(service, posDataAPIClient)
	app.MountTablepayController(service, c2)

	return &Server{
		service:  service,
		listener: listener,
		httpServer: &http.Server{
			Addr:    listener.Addr().String(),
			Handler: service.Mux,
		},
	}
}

// Run starts the server.
// Blocks until an error occurs or Shutdown is called.
func (s *Server) Run() error {
	if err := s.httpServer.Serve(s.listener); err != nil && err != http.ErrServerClosed {
		return errors.Wrap(err, "error running server")
	}
	return nil
}

// Shutdown stops the server.
func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutting down server")
	}
	return nil
}

// Addr returns the listening address of the server.
func (s *Server) Addr() string {
	return s.httpServer.Addr
}
