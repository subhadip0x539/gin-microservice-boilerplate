package logging

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitializeLogger() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	Logger = slog.New(logHandler)
}
