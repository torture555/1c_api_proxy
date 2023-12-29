package app

import (
	"1c_api_proxy/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func StartConnectionSQL() {

	fileConf := "config/sql.json"

	configModel := models.ConfSQL{}
	dataFile, err := os.ReadFile(fileConf)
	if err != nil {
		var raw models.Log
		raw.Context = err.Error()
		raw.Comment = "Запуск соединения с SQL"
		raw.Error("Не удалось прочитать или найти sql.json")
		panic(raw)
	}
	err = json.Unmarshal(dataFile, &configModel)
	if err != nil {
		var raw models.Log
		raw.Context = err.Error()
		raw.Comment = "Запуск соединения с SQL"
		raw.Error("Не удалось прочитать sql.json по формату JSON")
		panic(raw)
	}
	if configModel.TypeSQL != "mysql" {
		var raw models.Log
		raw.Context = "Тип СУБД указан " + configModel.TypeSQL
		raw.Comment = "Запуск соединения с SQL"
		raw.Error("Тип СУБД отличается от возможного mysql")
		panic(raw)
	}

	strAccess := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configModel.Login, configModel.Password, configModel.Host, configModel.Port, configModel.DBname)
	models.DB.DB, err = sql.Open(configModel.TypeSQL, strAccess)
	if err != nil {
		var raw models.Log
		raw.Context = err.Error()
		raw.Comment = "Запуск соединения с SQL"
		raw.Error("Не удалось подключится к базе данных MySQL проверьте конфиг или возможность подключения к СУБД")
		panic(raw)
	}

	if !models.DB.CheckSchema() {
		if !models.DB.InitSchema() {
			panic("Не удалось инициализировать схему БД")
		}
	}

}
