package data

import (
	"1c_api_proxy/internal/models"
	"errors"
)

func WriteInLog(log *models.Log) error {

	err := errors.New("Не удалось зафиксировать запись в журнал")
	return err

}
