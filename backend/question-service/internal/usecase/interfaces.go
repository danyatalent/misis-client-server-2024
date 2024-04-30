package usecase

import "github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/models"

type QuestionService interface {
	GetQuestion() (models.Question, error)
}
