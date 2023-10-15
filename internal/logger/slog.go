package logger

import (
	"log/slog"
	"os"
)

func NewSlog() *slog.Logger {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &opts))
}
