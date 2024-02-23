package app

import (
	"1c_api_proxy/internal/services/database"
	"errors"
	"log/slog"
	"os"
)

func StartLogger() {

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	}))
	slog.SetDefault(logger)

	err := errors.New("")
	database.FileLog, err = os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		database.FileLog = nil
	}

}
