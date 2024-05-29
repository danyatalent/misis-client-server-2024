package usecase

import (
	"context"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=WebAPI
type QuestionWebAPI interface {
	GetOneQuestion() (*models.Question, error)
	GetAllQuestions() ([]*models.Question, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=Cache
type QuestionCache interface {
	Put(ctx context.Context, value *models.Question) error
	GetRandom(ctx context.Context) (*models.Question, error)
}
