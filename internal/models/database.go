package models

import (
	"1c_api_proxy/internal/services/database"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"os"
)

var DBConnect DataBase

type DataBase struct {
	clickhouse.Conn
}

type DataBaseUse interface {
	AddLog(log *Log, level string) error
	CheckSchema() bool
	InitTable() bool
	IsConnected() bool
}

func (db *DataBase) AddLog(log *Log, level string) error {
	query := database.TemplateAddLog()
	query = fmt.Sprintf(query, log.BaseName, log.Context, log.Comment, log.Handler, level)
	_, err := db.Query(context.Background(), query)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) CheckSchema() bool {
	query := database.TemplateCheckSchema()
	_, err := db.Query(context.Background(), query)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (db *DataBase) InitTable() bool {
	query := database.TemplateInitSchema()
	_, err := db.Query(context.Background(), query)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (db *DataBase) IsConnected() bool {
	err := db.Ping(context.Background())
	if err != nil {
		return false
	}
	return true
}

func GetSettingsDB() (*ConfSQL, error) {

	fileConf := "config/database.json"

	configModel := ConfSQL{}
	dataFile, err := os.ReadFile(fileConf)
	if err != nil {
		var raw Log
		raw.Context = err.Error()
		raw.Comment = "Запуск соединения с Clickhouse"
		raw.Error("Не удалось прочитать или найти database.json")
	}
	err = json.Unmarshal(dataFile, &configModel)
	if err != nil {
		var raw Log
		raw.Context = err.Error()
		raw.Comment = "Запуск соединения с Clickhouse"
		raw.Error("Не удалось прочитать database.json по формату JSON")
	}

	return &configModel, err

}

func SetSettingsDB(model *ConfSQL) bool {

	databaseConf, err := os.OpenFile("config/database.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return false
	}
	defer databaseConf.Close()

	writer := bufio.NewWriter(databaseConf)
	_, err = writer.Write([]byte{})
	if err != nil {
		Log{
			Context: err.Error(),
			Comment: "Не удалось очистить файл database.json",
		}.Error("Не удалось очистить файл database.json")
		return false
	}

	newText, err := json.Marshal(*model)
	if err != nil {
		Log{
			Context: err.Error(),
			Comment: "Не удалось преобразовать настройки БД в формат JSON",
		}.Error("Не удалось преобразовать настройки БД в формат JSON")
	}

	_, err = writer.Write(newText)
	if err != nil {
		Log{
			Context: err.Error(),
			Comment: "Не удалось записать настройки БД в файл database.json",
		}.Error("Не удалось записать настройки БД в файл database.json")
		return false
	}

	return true

}
