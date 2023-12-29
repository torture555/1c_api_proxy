package app

import (
	"log/slog"
	"os"
)

func StartLogger() {

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	}))
	slog.SetDefault(logger)

}
