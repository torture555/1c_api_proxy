package models

import (
	"1c_api_proxy/internal/services/database"
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
)

var DBConnect DataBase

type DataBase struct {
	clickhouse.Conn
}

type DataBaseUse interface {
	AddLog(log *Log, level string) error
	CheckSchema() bool
	InitTable() bool
	IsConnected() bool
}

func (db *DataBase) AddLog(log *Log, level string) error {
	query := database.TemplateAddLog()
	query = fmt.Sprintf(query, log.BaseName, log.Context, log.Comment, log.Handler, level)
	_, err := db.Query(context.Background(), query)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) CheckSchema() bool {
	query := database.TemplateCheckSchema()
	_, err := db.Query(context.Background(), query)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (db *DataBase) InitTable() bool {
	query := database.TemplateInitSchema()
	_, err := db.Query(context.Background(), query)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (db *DataBase) IsConnected() bool {
	err := db.Ping(context.Background())
	if err != nil {
		return false
	}
	return true
}
