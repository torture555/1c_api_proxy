package models

import (
	"1c_api_proxy/internal/services/database"
	"database/sql"
	"fmt"
)

var DB DataBase

type DataBase struct {
	*sql.DB
}

func (db *DataBase) AddLog(log *Log, level string) error {
	query := database.TemplateAddLog()
	query = fmt.Sprintf(query, log.BaseID, log.BaseName, log.Context, log.InternalContext, log.Comment, level)

	err := DB.DB.Ping()
	_, err = db.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) CheckSchema() bool {
	query := database.TemplateCheckSchema()
	_, err := db.Query(query)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (db *DataBase) InitSchema() bool {
	query := database.TemplateInitSchema()
	_, err := db.Query(query)
	if err != nil {
		return false
	} else {
		return true
	}
}
