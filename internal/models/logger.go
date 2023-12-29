package models

import (
	"log/slog"
)

type Log struct {
	BaseID          int    `json:"baseID"`          // 1C base ID in service
	BaseName        string `json:"baseName"`        // 1C base name
	Context         string `json:"context"`         // context request
	InternalContext string `json:"internalContext"` // context service
	Comment         string `json:"comment"`         // text request
	Handler         string `json:"handler"`         // if api input handler
}

type LoggerProxy interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

func (l Log) Info(msg string) {
	if DB.DB != nil {
		_ = DB.AddLog(&l, "Info")
	}
	slog.Default().Info(msg+": ", "LogObj", l)
}

func (l Log) Warn(msg string) {
	if DB.DB != nil {
		_ = DB.AddLog(&l, "Warn")
	}
	slog.Default().Warn(msg+": ", "LogObj", l)
}

func (l Log) Error(msg string) {
	if DB.DB != nil {
		_ = DB.AddLog(&l, "Error")
	}
	slog.Default().Error(msg+": ", "LogObj", l)
}
