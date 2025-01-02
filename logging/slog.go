package logging

import (
	"fmt"
	"log/slog"
	"os"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
)

func init() {
	envLogLevel := os.Getenv("LOG_LEVEL")
	slogLevel := slog.LevelInfo

	switch envLogLevel {
	case LevelError:
		slogLevel = slog.LevelError
	case LevelWarn:
		slogLevel = slog.LevelWarn
	case LevelInfo:
		slogLevel = slog.LevelInfo
	case LevelDebug:
		slogLevel = slog.LevelDebug
	}

	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slogLevel,
	})
	slog.SetDefault(slog.New(textHandler))
	fmt.Println("==> log level(slog): " + slogLevel.Level().String())
}

