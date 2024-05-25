package v1

import (
	"context"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/pkg/logger"
	pb "github.com/danyatalent/protos/gen/go/answer"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type AnswerService interface {
	CheckAnswer(ctx context.Context, in *pb.AnswerRequest) (*pb.AnswerResponse, error)
}

type answerRoutes struct {
	l *slog.Logger
	t AnswerService
}

func NewAnswerRoutes(handler *gin.RouterGroup, t AnswerService, l *slog.Logger) {
	r := &answerRoutes{
		l: l,
		t: t,
	}
	h := handler.Group("/answer")
	{
		h.POST("/check", r.CheckAnswer)
	}
}

func (r *answerRoutes) CheckAnswer(c *gin.Context) {
	const op = "http/v1.answer/check"
	r.l = r.l.With(op)

	req := &pb.AnswerRequest{}

	if err := c.BindJSON(req); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	resp, err := r.t.CheckAnswer(c, req)
	if err != nil {
		r.l.Error("error checking answer", logger.Err(err))
		errorResponse(c, http.StatusInternalServerError, "error checking answer")

		return
	}
	c.JSON(http.StatusOK, resp)
}
