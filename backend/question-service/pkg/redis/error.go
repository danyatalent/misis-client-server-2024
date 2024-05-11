package redis

import (
	"errors"
	"github.com/redis/go-redis/v9"
)

var (
	ErrKeyNotExists = errors.New("key not exists")
	Nil             = redis.Nil
)
