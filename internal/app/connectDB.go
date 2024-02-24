package app

import (
	"1c_api_proxy/internal/models"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/go-sql-driver/mysql"
)

func StartConnectionSQL() {

	configModel, err := models.GetSettingsDB()
	if err != nil {
		return
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
		return
	}

	if !models.DBConnect.CheckSchema() {
		if !models.DBConnect.InitTable() {
			panic("Не удалось инициализировать таблицу БД")
		}
	}

}
