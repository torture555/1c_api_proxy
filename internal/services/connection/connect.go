package connection

import (
	"1c_api_proxy/internal/models"
	api_v1 "1c_api_proxy/internal/transport/rest/v1"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func ConnectLoop(thread *models.ThreadConnect1C) {
	statusFailedConnection := false
	secondsToRetry := models.RetryConnectSeconds * time.Second
	timerRetryConnect := time.NewTimer(secondsToRetry)
	chanStatusFailed := make(chan bool)

	for {
		if statusFailedConnection {
			countFailed := 0
			timerReconnect := time.NewTimer(models.RetryConnectSeconds * time.Second)
			for countFailed < models.ConstCountFailedConnections {
				<-timerReconnect.C
				timerReconnect.Reset(models.RetryConnectSeconds * time.Second)
				go pingConnect(thread.Client, thread.Base, chanStatusFailed)
				res := <-chanStatusFailed
				return
				if res {
					statusFailedConnection = false
					break
				}
				countFailed++
			}
			if statusFailedConnection {
				return
			}
		} else {
			select {
			case <-timerRetryConnect.C:
				timerRetryConnect.Reset(secondsToRetry)
				go pingConnect(thread.Client, thread.Base, chanStatusFailed)

			case res := <-chanStatusFailed:
				if !res {
					statusFailedConnection = true
				}

			default:

			}
		}

	}
}

func pingConnect(client *http.Client, base *models.Infobase, chanStatusFailed chan bool) {
	result := false
	request, err := http.NewRequest(api_v1.Path1C_PingConnection, base.URL+"/"+api_v1.Path1C_PingConnection, nil)
	if err != nil {
		go func(err error) {
			models.Log{
				BaseID:   base.Id,
				BaseName: base.Name,
				Context:  err.Error(),
				Handler:  api_v1.Path1C_PingConnection,
				Comment:  "Не удалось сформировать запрос Ping",
			}.Error("Не удалось сформировать запрос Ping")
		}(err)
		chanStatusFailed <- false
	}
	request.SetBasicAuth(base.Login, base.Password)

	response, err := client.Do(request)
	if err != nil {
		go func(err error) {
			models.Log{
				BaseID:          base.Id,
				BaseName:        base.Name,
				Context:         err.Error(),
				InternalContext: response.Status,
				Handler:         api_v1.Path1C_PingConnection,
				Comment:         "Не удалось получить ответ на запрос Ping",
			}.Error("Не удалось получить ответ на запрос Ping")
		}(err)
		chanStatusFailed <- false
	}

	if response.StatusCode == http.StatusOK {
		result = true
	} else {
		go func(status string) {
			models.Log{
				BaseID:   base.Id,
				BaseName: base.Name,
				Context:  status,
				Handler:  api_v1.Path1C_PingConnection,
				Comment:  "Получен неудачный ответ на Ping",
			}.Warn("Получен неудачный ответ на Ping")
		}(response.Status)
	}
	chanStatusFailed <- result
}

func ProxyRequest(thread *models.ThreadConnect1C, structChan *models.ModelChanConnect) {

	path := strings.ReplaceAll(structChan.C.Request.RequestURI, api_v1.PathProxy_Proxy, "")
	path = strings.ReplaceAll(path, "/"+thread.Base.Name, "")
	request, err := http.NewRequest(structChan.C.Request.Method, thread.Base.URL+path, structChan.C.Request.Body)
	if err != nil {
		comment := fmt.Sprintf("Метод: %v,\n Путь: %v,\n Заголовок: %v,\n Тело: %v,\n Источник: %v",
			structChan.C.Request.Method, path, structChan.C.Request.Header,
			structChan.C.Request.Body, structChan.C.Request.Host)
		go func(comment string, err error) {
			models.Log{
				BaseID:   thread.Base.Id,
				BaseName: thread.Base.Name,
				Context:  err.Error(),
				Comment:  comment,
			}.Error("Не удалось сформировать проксируемый запрос")
		}(comment, err)
		return
	}
	request.SetBasicAuth(thread.Base.Login, thread.Base.Password)
	request.Header = structChan.C.Request.Header.Clone()
	response, err := thread.Client.Do(request)
	if err != nil {
		comment := fmt.Sprintf("Метод: %v,\n Путь: %v,\n Заголовок: %v,\n Тело: %v,\n Источник: %v,\n Код ответа: %v",
			request.Method, path, request.Header, request.Body, structChan.C.Request.Host, response.StatusCode)
		go func(comment string) {
			models.Log{
				BaseID:   thread.Base.Id,
				BaseName: thread.Base.Name,
				Context:  err.Error(),
				Comment:  comment,
			}.Error("Не удалось получить ответ на проксируемый запрос")
		}(comment)
		return
	}
	body, err := io.ReadAll(response.Body)
	for i, v := range response.Header {
		if i != "Date" && i != "Content-Length" {
			structChan.C.Header(i, v[0])
		}
	}
	if body != nil {
		content := response.Header.Get("Content-Type")
		if content == "" {
			content = "application/json"
		}
		structChan.C.Data(response.StatusCode, content, body)
	} else {
		structChan.C.Status(response.StatusCode)
	}

	comment := fmt.Sprintf("Метод: %v,\n Путь: %v,\n Заголовок: %v,\n Тело: %v,\n Источник: %v,\n "+
		"Код ответа: %v\n Заголовок ответа: %v,\n Тело ответа: %v",
		request.Method, path, request.Header, request.Body, structChan.C.Request.Host, response.StatusCode,
		response.Header, string(body))
	go func(comment string) {
		models.Log{
			BaseID:   thread.Base.Id,
			BaseName: thread.Base.Name,
			Context:  "Получен проксируемый ответ",
			Comment:  comment,
		}.Info("Произведен проксируемый запрос")
	}(comment)

}
