package v1

import (
	"context"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/models"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type QuestionService interface {
	GetRandom(ctx context.Context) (*models.Question, error)
	Get(ctx context.Context, id string) (*models.Question, error)
}

type questionRoutes struct {
	t QuestionService
	l *slog.Logger
}

func NewQuestionRoutes(handler *gin.RouterGroup, t QuestionService, l *slog.Logger) {
	r := &questionRoutes{
		t: t,
		l: l,
	}
	h := handler.Group("/question")
	{
		h.GET("/random", r.random)
	}
}

type randomResponse struct {
	Question models.Question `json:"question"`
}

func (r *questionRoutes) random(c *gin.Context) {
	const op = "http/v1.question.random"
	r.l = r.l.With(op)

	q, err := r.t.GetRandom(c.Request.Context())
	if err != nil {
		r.l.Error("error getting random question", logger.Err(err))
		errorResponse(c, http.StatusInternalServerError, "internal server error")

		return
	}
	c.JSON(http.StatusOK, randomResponse{
		Question: *q,
	})
}
