// Package redis implements Redis connection.
package redis

import "time"

// Option -.
type Option func(*Redis)

// MaxRetries -.
func MaxRetries(retries int) Option {
	return func(r *Redis) {
		r.maxRetries = retries
	}
}

// DialTimeout -.
func DialTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.dialTimeout = timeout
	}
}

// ReadTimeout -.
func ReadTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.readTimeout = timeout
	}
}

// WriteTimeout -.
func WriteTimeout(timeout time.Duration) Option {
	return func(r *Redis) {
		r.writeTimeout = timeout
	}
}
