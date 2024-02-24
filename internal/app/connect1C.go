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

		connection.InitService1CAPI(&base)

	}

}
