package models

import (
	"1c_api_proxy/internal/transport/rest/front"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

var Connections connections1C

type connections1C struct {
	ThreadConnects map[string]ThreadConnect1C
}

type ThreadConnect1C struct {
	Base                *Infobase
	ChanResponseRequest chan ModelChanConnect
	Client              *http.Client
}

type ModelChanConnect struct {
	Check  bool
	Close  bool
	Result bool
	C      *gin.Context
}

func (threads *connections1C) FindThreadConnectByName(name string) *ThreadConnect1C {
	res := threads.ThreadConnects[name]

	if res.Base == nil {
		return nil
	}

	return &res
}

func (threads *connections1C) AddNewThread(base *Infobase) bool {
	if base.URL == "" || base.Name == "" {
		return false
	}
	newChan := make(chan ModelChanConnect)
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	thread := ThreadConnect1C{
		Base:                base,
		ChanResponseRequest: newChan,
		Client:              &client,
	}

	threads.ThreadConnects[base.Name] = thread

	go ReplaceConfigInfobases()

	return true
}

func (threads *connections1C) EditThread(base *Infobase) bool {

	findThread := threads.FindThreadConnectByName(base.Name)
	if findThread == nil {
		return false
	}

	findThread.Base = base

	go ReplaceConfigInfobases()

	return true
}

func (threads *connections1C) DeleteThread(thread *ThreadConnect1C) bool {
	thread.CloseLoop()
	close(thread.ChanResponseRequest)
	thread.Client.CloseIdleConnections()

	delete(threads.ThreadConnects, thread.Base.Name)

	go ReplaceConfigInfobases()

	return true
}

func (threads *connections1C) GetInfobasesList() *Infobases {

	list := Infobases{}

	for k, v := range threads.ThreadConnects {
		list.Bases = append(list.Bases, Infobase{
			Name:     k,
			URL:      v.Base.URL,
			Login:    v.Base.Login,
			Password: "",
		})
	}

	return &list

}

func (thread *ThreadConnect1C) ChanIsClosed() bool {
	model := makeEmptyModelChanConnect()
	model.Check = true
	thread.ChanResponseRequest <- *model
	res, ok := <-thread.ChanResponseRequest
	if !ok {
		return ok
	} else {
		if ok {
			return res.Close
		} else {
			return ok
		}
	}
}

func (thread *ThreadConnect1C) CloseLoop() {
	model := makeEmptyModelChanConnect()
	thread.ChanResponseRequest <- *model
}

func makeEmptyModelChanConnect() *ModelChanConnect {
	return &ModelChanConnect{
		Check:  false,
		Close:  true,
		C:      &gin.Context{},
		Result: false,
	}
}

func ReplaceConfigInfobases() bool {

	err := os.Truncate("config/infobases.json", 0)
	if err != nil {
		Log{
			Context: err.Error(),
			Comment: "Не удалось очистить хралищие данных подключения к БД",
			Handler: front.Database + "/" + front.SetDBParams,
		}.Error("Не удалось очистить хранилище данных подключения к БД")
		return false
	}

	databaseConf, err := os.OpenFile("config/infobases.json", os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return false
	}
	defer databaseConf.Close()

	model := Connections.GetInfobasesList()
	newText, err := json.Marshal(*model)
	if err != nil {
		writeEmptyParams()
		Log{
			Context: err.Error(),
			Comment: "Не удалось преобразовать настройки БД в формат JSON",
		}.Error("Не удалось преобразовать настройки БД в формат JSON")

	}

	_, err = databaseConf.Write(newText)
	if err != nil {
		writeEmptyParams()
		Log{
			Context: err.Error(),
			Comment: "Не удалось записать настройки БД в файл database.json",
		}.Error("Не удалось записать настройки БД в файл database.json")
		return false
	}

	return true

}
