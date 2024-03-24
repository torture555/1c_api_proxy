package models

import (
	"1c_api_proxy/internal/services/database"
	"1c_api_proxy/internal/transport/rest/front"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var DBConnect DataBase

type DataBase struct {
	*sql.DB
}

type Logs struct {
	L []RawLog `json:"l"`
}

type RawLog struct {
	Date     string `json:"date"`
	BaseName string `json:"base_name"`
	Context  string `json:"context"`
	Handler  string `json:"handler"`
	Comment  string `json:"comment"`
	Level    string `json:"level"`
}

type DataBaseUse interface {
	AddLog(log *Log, level string) error
	CheckSchema() bool
	InitTable() bool
	IsConnected() bool
}

func InitDB() error {

	configModel, err := GetSettingsDB()
	if err != nil {
		return err
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configModel.Host, configModel.Port, configModel.Login, configModel.Password, configModel.DBName)

	DBConnect.DB, err = sql.Open("postgres", connStr)

	if err != nil {
		var raw Log
		raw.Context = err.Error()
		raw.Comment = "Запуск соединения с Postgres"
		raw.Error("Не удалось подключится к базе данных Postgres проверьте конфиг или возможность подключения к СУБД")
		return err
	}

	if !DBConnect.CheckSchema() {
		if !DBConnect.InitTable() {
			panic("Не удалось инициализировать таблицу БД")
		}
	}

	return nil
}

func (db *DataBase) AddLog(log *Log, level string) error {
	if !db.IsConnected() {
		return errors.New("Не подключена БД")
	}
	query := database.TemplateAddLog()
	query = fmt.Sprintf(query, log.BaseName, log.Context, log.Comment, log.Handler, level)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) CheckSchema() bool {
	if !db.IsConnected() {
		return true
	}

	query := database.TemplateCheckSchema()
	_, err := db.Exec(query)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (db *DataBase) InitTable() bool {
	query := database.TemplateInitSchema()
	_, err := db.Exec(query)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (db *DataBase) IsConnected() bool {
	if db.DB == nil {
		return false
	}
	_, err := db.Exec(database.TemplateNow())
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

	err := os.Truncate("config/database.json", 0)
	if err != nil {
		Log{
			Context: err.Error(),
			Comment: "Не удалось очистить хралищие данных подключения к БД",
			Handler: front.Database + "/" + front.SetDBParams,
		}.Error("Не удалось очистить хранилище данных подключения к БД")
		return false
	}

	databaseConf, err := os.OpenFile("config/database.json", os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		writeEmptyParams()
		return false
	}
	defer databaseConf.Close()

	newText, err := json.Marshal(*model)
	if err != nil {
		writeEmptyParams()
		Log{
			Context: err.Error(),
			Comment: "Не удалось преобразовать настройки БД в формат JSON",
		}.Error("Не удалось преобразовать настройки БД в формат JSON")

	}

	_, err = databaseConf.Write(newText)
	if err != nil {
		writeEmptyParams()
		Log{
			Context: err.Error(),
			Comment: "Не удалось записать настройки БД в файл database.json",
		}.Error("Не удалось записать настройки БД в файл database.json")
		return false
	}

	_ = InitDB()

	return true

}

func writeEmptyParams() {

	databaseConf, err := os.OpenFile("config/database.json", os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return
	}
	defer databaseConf.Close()

	model := ConfSQL{}
	newText, _ := json.Marshal(model)
	_, err = databaseConf.Write(newText)
	if err != nil {
		Log{
			Context: err.Error(),
			Comment: "Не удалось записать настройки по умолчанию БД в файл database.json",
		}.Error("Не удалось записать настройки по умолчанию БД в файл database.json")
	}
}

func GetLogs() Logs {
	query := database.GetLogs()
	res, err := DBConnect.Query(query)
	if err != nil {
		Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать логи",
			Handler: front.Logs + "/" + front.GetLog,
		}.Error("Не удалось прочитать логи")
	}

	slice := make([]RawLog, 0)

	for res.Next() {
		newRaw := RawLog{}
		_ = res.Scan(&newRaw.Date, &newRaw.BaseName, &newRaw.Context, &newRaw.Comment, &newRaw.Handler, &newRaw.Level)
		slice = append(slice, newRaw)
	}
	resSlice := Logs{L: slice}
	return resSlice

}
