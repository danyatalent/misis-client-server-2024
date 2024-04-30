package app

import (
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/config"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/usecase/webAPI"
)

func Run(cfg *config.Config) error {
	//log := logger.InitLogger(cfg.Env)
	_, _ = webAPI.GetQuestion()
	//log.Info().Any("quest", quest).Msg("quest")
	return nil
}
