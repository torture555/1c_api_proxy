package models

import (
	"1c_api_proxy/internal/services/database"
	"fmt"
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
	if DBConnect.IsConnected() {
		_ = DBConnect.AddLog(&l, "Info")
	}

	if database.FileLog != nil {
		database.AddStr(fmt.Sprintf("%s: %+v", "INFO", l))
	}

	slog.Default().Info(msg+": ", "LogObj", l)
}

func (l Log) Warn(msg string) {
	if DBConnect.IsConnected() {
		_ = DBConnect.AddLog(&l, "Warn")
	}

	if database.FileLog != nil {
		database.AddStr(fmt.Sprintf("%s: %+v", "WARN: ", l))
	}

	slog.Default().Warn(msg+": ", "LogObj", l)
}

func (l Log) Error(msg string) {
	if DBConnect.IsConnected() {
		_ = DBConnect.AddLog(&l, "Error")
	}

	if database.FileLog != nil {
		database.AddStr(fmt.Sprintf("%s: %+v", "ERROR: ", l))
	}

	slog.Default().Error(msg+": ", "LogObj", l)
}
