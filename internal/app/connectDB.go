package app

import (
	"1c_api_proxy/internal/models"
	"encoding/json"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func StartConnectionSQL() {

	fileConf := "config/database.json"

	configModel := models.ConfSQL{}
	dataFile, err := os.ReadFile(fileConf)
	if err != nil {
		var raw models.Log
		raw.Context = err.Error()
		raw.Comment = "Запуск соединения с Clickhouse"
		raw.Error("Не удалось прочитать или найти database.json")
		panic(raw)
	}
	err = json.Unmarshal(dataFile, &configModel)
	if err != nil {
		var raw models.Log
		raw.Context = err.Error()
		raw.Comment = "Запуск соединения с Clickhouse"
		raw.Error("Не удалось прочитать database.json по формату JSON")
		panic(raw)
	}

	models.DBConnect.Conn, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", configModel.Host, configModel.Port)},
		Auth: clickhouse.Auth{
			Username: configModel.Login,
			Password: configModel.Password,
		},
	})
	if err != nil {
		var raw models.Log
		raw.Context = err.Error()
		raw.Comment = "Запуск соединения с Clickhouse"
		raw.Error("Не удалось подключится к базе данных Clickhouse проверьте конфиг или возможность подключения к СУБД")
		panic(raw)
	}

	if !models.DBConnect.CheckSchema() {
		if !models.DBConnect.InitTable() {
			panic("Не удалось инициализировать таблицу БД")
		}
	}

}
