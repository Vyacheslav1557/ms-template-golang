package logger

import (
	"fmt"
	"log/slog"
	"os"
)

var log *slog.Logger

func SetupLogger(env string) {
	switch env {
	case "dev":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		slog.Error(fmt.Sprintf(`env: "prod" or "dev" expected, got "%s"`, env))
	}
}

func Logger() slog.Logger {
	if log == nil {
		slog.Error("logger was not initialized")
		os.Exit(1)
	}
	return *log
}
