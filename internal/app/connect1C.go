package app

import (
	"1c_api_proxy/internal/models"
	"1c_api_proxy/internal/services/connection"
	"encoding/json"
	"log"
	"os"
)

func StartServices1CAPI() {

	basesModel := models.Infobases{}
	pathBases := "config/infobases.json"
	dataFile, err := os.ReadFile(pathBases)
	if err != nil {
		log.Fatal("Не удалось прочитать файл конфига infobases.json" + err.Error())
	}

	err = json.Unmarshal(dataFile, &basesModel)
	if err != nil {
		log.Fatal("Не удалось расшифровать JSON infobases.json" + err.Error())
	}

	for _, base := range basesModel.Bases {

		initService1CAPI(&base)

	}

}

func initService1CAPI(base *models.Infobase) {

	res := models.Connections.AddNewThread(base)

	if res {
		go connection.ConnectLoop(models.Connections.FindThreadConnectByName(base.Name))
		models.Log{
			BaseID:   base.Id,
			BaseName: base.Name,
		}.Info("Инициализировано соеденине с информационной базой")
	} else {
		models.Log{
			BaseID:   base.Id,
			BaseName: base.Name,
		}.Warn("Не удалось инициализировать соединение с информационной базой")
	}

}
