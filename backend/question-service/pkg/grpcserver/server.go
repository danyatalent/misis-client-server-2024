// Package grpcserver implements gRPC server.
package grpcserver

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	_defaultAddr            = ":50051"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *grpc.Server
	addr            string
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(gRPCServer *grpc.Server, opts ...Option) *Server {
	s := &Server{
		server:          gRPCServer,
		addr:            _defaultAddr,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		listener, err := net.Listen("tcp", s.addr)
		if err != nil {
			s.notify <- err
			close(s.notify)
			return
		}
		s.notify <- s.server.Serve(listener)
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	ch := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(ch)
	}()
	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		s.server.Stop()
		return ctx.Err()
	}
}

func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.server.RegisterService(desc, impl)
}
