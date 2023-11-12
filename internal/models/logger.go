package models

import "time"

type Log struct {
	Datetime        string `json:"datetime"`        // "2023-11-11 21:55:55"
	BaseID          int    `json:"baseID"`          // 1C base ID in service
	BaseName        string `json:"baseName"`        // 1C base name
	Context         string `json:"context"`         // context request
	InternalContext string `json:"internalContext"` // context service
	Comment         string `json:"comment"`         // text request
	Handler         string `json:"handler"`         // if api input handler
}

func (log *Log) SetDataTime() {

	log.Datetime = time.DateTime

}
