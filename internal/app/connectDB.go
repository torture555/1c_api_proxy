package app

import (
	"1c_api_proxy/internal/models"
	_ "github.com/lib/pq"
)

func StartConnectionSQL() {

	err := models.InitDB()
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось инициализировать подключение к БД",
		}.Error("Не удалось инициализировать подключение к БД")
	}

}
