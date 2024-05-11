package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	_defaultMaxRetries   = 3
	_defaultDialTimeout  = 5 * time.Second
	_defaultReadTimeout  = 3 * time.Second
	_defaultWriteTimeout = 3 * time.Second
)

// Redis -.
type Redis struct {
	maxRetries   int
	dialTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration

	Client *redis.Client
}

// New -.
func New(addr string, opts ...Option) (*Redis, error) {
	r := &Redis{
		maxRetries:   _defaultMaxRetries,
		dialTimeout:  _defaultDialTimeout,
		readTimeout:  _defaultReadTimeout,
		writeTimeout: _defaultWriteTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(r)
	}

	options := &redis.Options{
		Addr:         addr,
		MaxRetries:   r.maxRetries,
		DialTimeout:  r.dialTimeout,
		ReadTimeout:  r.readTimeout,
		WriteTimeout: r.writeTimeout,
	}

	r.Client = redis.NewClient(options)

	ctx := context.Background()
	if err := r.Client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis - NewRedis - Ping: %w", err)
	}

	return r, nil
}

// Close -.
func (r *Redis) Close() {
	if r.Client != nil {
		r.Client.Close()
	}
}
