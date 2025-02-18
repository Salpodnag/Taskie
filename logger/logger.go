package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	l := slog.New(handler)

	return l
}
