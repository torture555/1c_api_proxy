package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Connections connections1C

type connections1C struct {
	ThreadConnects []ThreadConnect1C
}

type ThreadConnect1C struct {
	Base                *Infobase
	ChanResponseRequest chan ModelChanConnect
	Client              *http.Client
}

type ModelChanConnect struct {
	check  bool
	close  bool
	Result bool
	C      *gin.Context
}

func (threads *connections1C) FindThreadConnectByName(name string) *ThreadConnect1C {
	for _, val := range threads.ThreadConnects {
		if val.Base.Name == name {
			return &val
		}
	}
	return nil
}

func (threads *connections1C) FindIndexThreadByObj(thread *ThreadConnect1C) (bool, int) {
	for i, val := range threads.ThreadConnects {
		if &val == thread {
			return true, i
		}
	}
	return false, 0
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
	threads.ThreadConnects = append(threads.ThreadConnects, thread)

	return true
}

func (threads *connections1C) DeleteThread(thread *ThreadConnect1C) bool {
	closeChan := makeEmptyModelChanConnect()
	closeChan.close = true
	thread.ChanResponseRequest <- *closeChan
	close(thread.ChanResponseRequest)
	thread.Client.CloseIdleConnections()
	resFind, index := threads.FindIndexThreadByObj(thread)
	if resFind {
		threads.ThreadConnects = append(threads.ThreadConnects[:index], threads.ThreadConnects[index+1:]...)
	}
	return true
}

func (chanConnect *ModelChanConnect) ChanClosed() bool {
	return chanConnect.close
}

func makeEmptyModelChanConnect() *ModelChanConnect {
	return &ModelChanConnect{
		close:  false,
		C:      &gin.Context{},
		Result: false,
	}
}
