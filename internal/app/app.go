package app

import (
	"1c_api_proxy/internal/models"
	"encoding/json"
	"log"
	"os"
)

func StartServices1CAPI() {

	exPath, _ := os.Executable()
	configPath := exPath + "/config/"
	fileConf := configPath + "configs.json"

	configModel := models.Service1CAPI{}
	dataFile, err := os.ReadFile(fileConf)
	if err != nil {
		log.Fatal("Не удалось прочитать или найти configs.json" + err.Error())
	}

	err = json.Unmarshal(dataFile, &configModel)
	if err != nil {
		log.Fatal("Не удалось прочитать конфиг файл configs.json." +
			"Проверьте правильность написания конфига" + err.Error())
	}

	basesModel := models.Infobases{}
	pathBases := configPath + "infobases.json"
	dataFile, err = os.ReadFile(pathBases)
	if err != nil {
		log.Fatal("Не удалось прочитать файл конфига infobases.json" + err.Error())
	}

	err = json.Unmarshal(dataFile, &basesModel)
	if err != nil {
		log.Fatal("Не удалось расшифровать JSON infobases.json" + err.Error())
	}

	for _, base := range basesModel.Bases {

		statusChan := make(chan bool)

		go initService1CAPI(&base, statusChan)

		result := <-statusChan
		if result == false {
			log.Panic("Не получилось инициализировать подключение в инф.базе " + base.Name)
		} else {

		}

	}

}

func initService1CAPI(base *models.Infobase, status chan bool) {

}
