package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server represents an HTTP server wrapper with graceful shutdown capabilities
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New creates a new HTTP server instance with the given handler and options
func New(handler http.Handler, opts ...Option) *Server {
	if handler == nil {
		panic("handler cannot be nil")
	}

	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Apply all provided options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

// start begins listening for incoming HTTP requests in a separate goroutine
func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify returns a channel that will receive server errors if any occur
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown gracefully shuts down the server without interrupting active connections
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}