package grpcserver

import (
	"time"

	"google.golang.org/grpc"
)

// Option -.
type Option func(*Server)

// Addr -.
func Addr(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

// UnaryInterceptor -.
func UnaryInterceptor(interceptor grpc.UnaryServerInterceptor) Option {
	return func(s *Server) {
		s.server = grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	}
}

// StreamInterceptor -.
func StreamInterceptor(interceptor grpc.StreamServerInterceptor) Option {
	return func(s *Server) {
		s.server = grpc.NewServer(grpc.StreamInterceptor(interceptor))
	}
}
