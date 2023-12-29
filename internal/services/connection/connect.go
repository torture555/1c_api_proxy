package connection

import (
	"1c_api_proxy/internal/models"
	api_v1 "1c_api_proxy/internal/transport/rest/v1"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func ConnectLoop(thread *models.ThreadConnect1C) {
	statusFailedConnection := false
	secondsToRetry := models.RetryConnectSeconds * time.Second
	timerRetryConnect := time.NewTimer(secondsToRetry)
	client := http.Client{}
	chanStatusFailed := make(chan bool)

	for {
		if statusFailedConnection {
			countFailed := 0
			timerReconnect := time.NewTimer(1 * time.Second)
			for countFailed < models.ConstCountFailedConnections {
				<-timerReconnect.C
				timerReconnect.Reset(1 * time.Second)
				pingConnect(&client, thread.Base, chanStatusFailed)
				res := <-chanStatusFailed
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
			case structChan := <-thread.ChanResponseRequest:
				if structChan.ChanClosed() {
					timerRetryConnect.Stop()
					return
				}
				go proxyRequest(&client, thread, &structChan)

			case <-timerRetryConnect.C:
				go pingConnect(&client, thread.Base, chanStatusFailed)

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
	request, err := http.NewRequest(api_v1.Path1C_PingConnection, base.URL, nil)
	if err != nil {
		models.Log{
			BaseID:   base.Id,
			BaseName: base.Name,
			Context:  err.Error(),
			Handler:  api_v1.Path1C_PingConnection,
		}.Error("Не удалось сформировать запрос Ping")
		chanStatusFailed <- false
	}
	request.SetBasicAuth(base.Login, base.Password)

	response, err := client.Do(request)
	if err != nil {
		models.Log{
			BaseID:          base.Id,
			BaseName:        base.Name,
			Context:         err.Error(),
			InternalContext: response.Status,
			Handler:         api_v1.Path1C_PingConnection,
		}.Error("Не удалось получить ответ на запрос Ping")
		chanStatusFailed <- false
	}

	if response.StatusCode == http.StatusOK {
		result = true
	} else {
		models.Log{
			BaseID:   base.Id,
			BaseName: base.Name,
			Context:  response.Status,
			Handler:  api_v1.Path1C_PingConnection,
		}.Warn("Получен неудачный ответ на Ping")
	}
	chanStatusFailed <- result
}

func proxyRequest(client *http.Client, chanConnect *models.ThreadConnect1C, structChan *models.ModelChanConnect) {

	path := strings.ReplaceAll(structChan.Request.RequestURI, api_v1.PathProxy_Proxy, "")

	request, err := http.NewRequest(structChan.Request.Method, chanConnect.Base.URL+path, structChan.Request.Body)
	if err != nil {
		models.Log{
			BaseID:   chanConnect.Base.Id,
			BaseName: chanConnect.Base.Name,
			Context:  err.Error(),
			InternalContext: fmt.Sprintf("Метод: %v,\n Путь: %v,\n Заголовок: %v,\n Тело: %v,\n Источник: %v",
				structChan.Request.Method, path, structChan.Request.Header,
				structChan.Request.Body, structChan.Request.Host),
		}.Error("Не удалось сформировать проксируемый запрос")
	}
	//request.SetBasicAuth(chanConnect.Base.Login, chanConnect.Base.Password)

	response, err := client.Do(request)
	if err != nil {
		models.Log{
			BaseID:   chanConnect.Base.Id,
			BaseName: chanConnect.Base.Name,
			Context:  err.Error(),
			InternalContext: fmt.Sprintf("Метод: %v,\n Путь: %v,\n Заголовок: %v,\n Тело: %v,\n Источник: %v,\n Код ответа: %v",
				request.Method, path, request.Header, request.Body, structChan.Request.Host, response.StatusCode),
		}.Error("Не удалось получить ответ на проксируемый запрос")
	}

	structChan.Response = *response
	structChan.Result = err == nil
	chanConnect.ChanResponseRequest <- *structChan
	models.Log{
		BaseID:   chanConnect.Base.Id,
		BaseName: chanConnect.Base.Name,
		Context:  "Получен проксируемый ответ",
		InternalContext: fmt.Sprintf("Метод: %v,\n Путь: %v,\n Заголовок: %v,\n Тело: %v,\n Источник: %v,\n "+
			"Код ответа: %v\n Заголовок ответа: %v,\n Тело ответа: %v",
			request.Method, path, request.Header, request.Body, structChan.Request.Host, response.StatusCode,
			response.Header, response.Body),
	}.Info("Произведен проксируемый запрос")

}
