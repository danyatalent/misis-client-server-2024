package logger

import (
	"github.com/rs/zerolog"
	"os"
)

const (
	envProd        = "prod"
	envDev         = "dev"
	envLocal       = "local"
	skipFrameCount = 3
)

// Logger -.
type Logger struct {
	*zerolog.Logger
}

func InitLogger(env string) *zerolog.Logger {
	switch env {
	case envLocal:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		out := zerolog.ConsoleWriter{Out: os.Stdout}
		logger := zerolog.New(out).With().Timestamp().Logger()
		return &logger
	case envDev:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()
		return &logger
	case envProd:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()
		return &logger
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		return &logger
	}
}
