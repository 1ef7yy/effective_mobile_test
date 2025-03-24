package logger

import (
	"log/slog"
	"os"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

type logger struct {
	log *slog.Logger
}

func NewLogger() Logger {
	opts := slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug.Level(),
	}
	log := slog.NewJSONHandler(os.Stdout, &opts)

	return logger{
		log: slog.New(log),
	}

}
