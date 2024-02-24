package models

import (
	"bufio"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
	newChan := make(chan ModelChanConnect, 100)
	client := http.Client{}
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
		return res.Close || ok
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

	databaseConf, err := os.OpenFile("config/infobases.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return false
	}
	defer databaseConf.Close()

	writer := bufio.NewWriter(databaseConf)
	_, err = writer.Write([]byte{})
	if err != nil {
		Log{
			Context: err.Error(),
			Comment: "Не удалось очистить файл infobases.json",
		}.Error("Не удалось очистить файл infobases.json")
		return false
	}

	model := Connections.GetInfobasesList()

	newText, err := json.Marshal(*model)
	if err != nil {
		Log{
			Context: err.Error(),
			Comment: "Не удалось преобразовать список баз в JSON",
		}.Error("Не удалось преобразовать список баз в JSON")
	}

	_, err = writer.Write(newText)
	if err != nil {
		Log{
			Context: err.Error(),
			Comment: "Не удалось записать новый список в файл infobases.json",
		}.Error("Не удалось записать новый список в файл infobases.json")
		return false
	}

	return true

}
