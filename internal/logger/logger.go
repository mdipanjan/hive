package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func init() {
	level := slog.LevelError // default: only errors

	if os.Getenv("DEBUG") == "1" {
		level = slog.LevelDebug
	}

	Log = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))
}
