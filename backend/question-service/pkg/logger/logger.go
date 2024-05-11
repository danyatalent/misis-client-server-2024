package logger

import (
	"github.com/phsym/console-slog"
	"log/slog"
	"os"
	"time"
)

const (
	envProd  = "prod"
	envDev   = "dev"
	envLocal = "local"
)

func InitLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envLocal:
		logger = slog.New(
			console.NewHandler(os.Stdout, &console.HandlerOptions{
				Level:      slog.LevelDebug,
				TimeFormat: time.TimeOnly,
			}),
		)
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return logger
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
