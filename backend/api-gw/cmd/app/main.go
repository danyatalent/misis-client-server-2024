package main

import (
	"github.com/danyatalent/misis-client-server-2024/backend/api-gw/config"
	"github.com/danyatalent/misis-client-server-2024/backend/api-gw/internal/app"
)

func main() {
	cfg := config.GetConfig()

	app.Run(cfg)
}
