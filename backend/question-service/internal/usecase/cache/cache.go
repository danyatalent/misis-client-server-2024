package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/apperror"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/models"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/pkg/logger"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/pkg/redis"
	"log/slog"
	"sync"
)

const (
	cachePrefix = "question"
)

type QuestionCache struct {
	logger *slog.Logger
	*redis.Redis
	mu sync.Mutex
}

func (c *QuestionCache) Put(ctx context.Context, id string, value *models.Question) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Enhancing logger
	const op = "QuestionCache.Put"
	c.logger = c.logger.With(op)

	// Marshall to JSON
	encodedQuestion, err := json.Marshal(&value)
	if err != nil {
		c.logger.Error("error marshalling to json", logger.Err(err))
		return fmt.Errorf("error marshalling to json: %w", err)
	}

	key := fmt.Sprintf("%s:%s", cachePrefix, id)

	// Using our cache struct to save
	err = c.Client.Set(ctx, key, encodedQuestion, 0).Err()
	if err != nil {
		c.logger.Error("error put in cache", logger.Err(err))
		return fmt.Errorf("error put in cache: %w", err)
	}
	return nil
}

func (c *QuestionCache) Get(ctx context.Context, id string) (*models.Question, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Enhancing logger
	const op = "QuestionCache.Get"
	c.logger = c.logger.With(op)

	key := fmt.Sprintf("%s:%s", cachePrefix, id)

	// Using our cache to get the model
	encodedQuestion, err := c.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		c.logger.Error("error key not exists in cache", logger.Err(err))
		return nil, apperror.ErrQuestionNotFound
	} else if err != nil {
		c.logger.Error("error getting from redis", logger.Err(err))
		return nil, fmt.Errorf("error getting from redis: %w", err)
	}

	question := &models.Question{}

	// Unmarshalling from JSON
	err = json.Unmarshal([]byte(encodedQuestion), question)
	if err != nil {
		c.logger.Error("error unmarshalling from json", logger.Err(err))
		return nil, fmt.Errorf("error unmarshalling from json: %w", err)
	}

	return question, nil
}

func (c *QuestionCache) GetRandom(ctx context.Context) (*models.Question, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	const op = "QuestionCache.GetRandom"
	c.logger = c.logger.With(op)

	randomKey, err := c.Client.RandomKey(ctx).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, apperror.ErrQuestionNotFound
		}
		c.logger.Error("error getting random key from redis", logger.Err(err))
		return nil, fmt.Errorf("error getting random key from redis: %w", err)
	}
	question := &models.Question{}

	rawJSON, err := c.Client.Get(ctx, randomKey).Result()
	if err != nil {
		c.logger.Error("error getting from redis", logger.Err(err))
		return nil, fmt.Errorf("error getting from redis: %w", err)
	}
	err = json.Unmarshal([]byte(rawJSON), question)
	if err != nil {
		c.logger.Error("error unmarshalling from json", logger.Err(err))
		return nil, fmt.Errorf("error unmarshalling from json: %w", err)
	}
	return question, nil

}

func New(logger *slog.Logger, redis *redis.Redis) *QuestionCache {
	return &QuestionCache{
		logger: logger,
		Redis:  redis,
	}
}
