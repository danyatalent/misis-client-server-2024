package main

import (
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/config"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/app"
)

func main() {
	cfg := config.GetConfig()
	err := app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
