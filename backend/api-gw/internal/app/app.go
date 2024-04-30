package app

import (
	"github.com/danyatalent/misis-client-server-2024/backend/api-gw/config"
	"github.com/danyatalent/misis-client-server-2024/backend/api-gw/pkg/logger"
)

func Run(cfg *config.Config) {
	log := logger.InitLogger(cfg.Env)
	log.Info().Msg("hello")

}
