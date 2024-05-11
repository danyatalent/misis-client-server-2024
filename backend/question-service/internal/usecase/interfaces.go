package usecase

import (
	"context"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/models"
)

type QuestionWebAPI interface {
	GetOneQuestion() (*models.Question, error)
	GetAllQuestions() ([]*models.Question, error)
}

type QuestionCache interface {
	Put(ctx context.Context, key string, value *models.Question) error
	Get(ctx context.Context, key string) (*models.Question, error)
	GetRandom(ctx context.Context) (*models.Question, error)
}
