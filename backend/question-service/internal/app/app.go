package app

import (
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/config"
	v1 "github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/controller/http/v1"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/usecase"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/usecase/cache"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/usecase/webAPI"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/pkg/httpserver"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/pkg/logger"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/pkg/redis"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) error {
	l := logger.InitLogger(cfg.Env)
	rdb, err := redis.New(cfg.Redis.Addr)
	if err != nil {
		l.Error("error connect to redis", logger.Err(err))
		return err
	}
	defer rdb.Close()

	questionUseCase := usecase.New(
		l,
		webAPI.New(l),
		cache.New(l, rdb),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, questionUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	l.Info("starting http server on port " + cfg.HTTP.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error("httpServer.Notify() failed", logger.Err(err))
	}
	err = httpServer.Shutdown()
	if err != nil {
		l.Error("server shutdown failed", logger.Err(err))
		return err
	}
	return nil
}
