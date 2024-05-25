package v1

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"log/slog"
	"net/http"
)

func NewRouter(handler *gin.Engine, l *slog.Logger, t QuestionService, a AnswerService) {
	//handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(cors.Default())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	h := handler.Group("/api/v1")
	{
		NewQuestionRoutes(h, t, l)
		NewAnswerRoutes(h, a, l)
	}
}
