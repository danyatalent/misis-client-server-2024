package usecase

import (
	"context"
	"errors"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/apperror"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/models"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/pkg/logger"
	"log/slog"
)

type QuestionUseCase struct {
	logger *slog.Logger
	webAPI QuestionWebAPI
	cache  QuestionCache
}

func NewQuestion(l *slog.Logger, w QuestionWebAPI, c QuestionCache) *QuestionUseCase {
	return &QuestionUseCase{
		logger: l,
		webAPI: w,
		cache:  c,
	}
}

//func (u *QuestionUseCase) Get(ctx context.Context, id string) (*models.Question, error) {
//	const op = "QuestionUseCase.Get"
//	u.logger.With(op)
//
//	question, err := u.cache.Get(ctx, id)
//	if err != nil {
//		if !errors.As(err, &apperror.ErrQuestionNotFound) {
//			u.logger.Error("unexpected error while getting data from cache", logger.Err(err))
//			return nil, err
//		}
//		questions, err := u.webAPI.GetAllQuestions()
//		if err != nil {
//			u.logger.Error("unexpected error while getting data from WebAPI", logger.Err(err))
//			return nil, err
//		}
//
//		for _, q := range questions {
//			err = u.cache.Put(ctx, q.ID, q)
//			if err != nil {
//				u.logger.Error("unexpected error while putting data in cache", logger.Err(err))
//			}
//
//		}
//		if len(questions) > 0 {
//			question = questions[0]
//		}
//	}
//	return question, nil
//}

func (u *QuestionUseCase) GetRandom(ctx context.Context) (*models.Question, error) {
	const op = "QuestionUseCase.GetRandom"
	u.logger.With(op)
	question, err := u.cache.GetRandom(ctx)
	if err != nil {
		if !errors.As(err, &apperror.ErrQuestionNotFound) {
			u.logger.Error("unexpected error while getting data from cache", logger.Err(err))
			return nil, err
		}
		questions, err := u.webAPI.GetAllQuestions()
		if err != nil {
			u.logger.Error("unexpected error while getting data from WebAPI", logger.Err(err))
			return nil, err
		}
		for i := 1; i < len(questions); i++ {
			err = u.cache.Put(ctx, questions[i])
			if err != nil {
				u.logger.Error("unexpected error while putting data in cache", logger.Err(err))
			}
		}
		if len(questions) > 0 {
			question = questions[0]
		}
	}
	return question, nil
}
