package data

import (
	"1c_api_proxy/internal/models"
	"errors"
	"reflect"
)

func AddNewInfobase(base *models.Infobase) error {

	err := errors.New("Не удалось добавить новую информационную базу в список")
	return err

}

func GetInfobaseByNameId(id, name interface{}) (models.Infobase, error) {

	presentID := reflect.ValueOf(id).String()
	presentName := reflect.ValueOf(name).String()
	err := errors.New("Не удалось получить информационную базу ID: " + presentID + " Name: " + presentName)
	if id == nil && name == nil {
		return models.Infobase{}, err
	}
	if reflect.TypeOf(id).Name() != "int" && reflect.TypeOf(name).Name() != "string" {
		return models.Infobase{}, err
	}

	return models.Infobase{}, err

}

func GetServiceInfobase(base *models.Infobase) models.ServiceInfobase1C {

	return models.ServiceInfobase1C{}

}
