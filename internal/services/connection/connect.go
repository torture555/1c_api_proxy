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
	chanStatusFailed := make(chan bool)

	for {
		if statusFailedConnection {
			countFailed := 0
			timerReconnect := time.NewTimer(1 * time.Second)
			for countFailed < models.ConstCountFailedConnections {
				<-timerReconnect.C
				timerReconnect.Reset(1 * time.Second)
				pingConnect(thread.Client, thread.Base, chanStatusFailed)
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
	//request.SetBasicAuth(base.Login, base.Password)

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

func ProxyRequest(thread *models.ThreadConnect1C, structChan *models.ModelChanConnect) {

	path := strings.ReplaceAll(structChan.C.Request.RequestURI, api_v1.PathProxy_Proxy, "")

	request, err := http.NewRequest(structChan.C.Request.Method, thread.Base.URL+path, structChan.C.Request.Body)
	if err != nil {
		models.Log{
			BaseID:   thread.Base.Id,
			BaseName: thread.Base.Name,
			Context:  err.Error(),
			InternalContext: fmt.Sprintf("Метод: %v,\n Путь: %v,\n Заголовок: %v,\n Тело: %v,\n Источник: %v",
				structChan.C.Request.Method, path, structChan.C.Request.Header,
				structChan.C.Request.Body, structChan.C.Request.Host),
		}.Error("Не удалось сформировать проксируемый запрос")
	}
	//request.SetBasicAuth(chanConnect.Base.Login, chanConnect.Base.Password)

	response, err := thread.Client.Do(request)
	if err != nil {
		models.Log{
			BaseID:   thread.Base.Id,
			BaseName: thread.Base.Name,
			Context:  err.Error(),
			InternalContext: fmt.Sprintf("Метод: %v,\n Путь: %v,\n Заголовок: %v,\n Тело: %v,\n Источник: %v,\n Код ответа: %v",
				request.Method, path, request.Header, request.Body, structChan.C.Request.Host, response.StatusCode),
		}.Error("Не удалось получить ответ на проксируемый запрос")
	}
	var body []byte
	header := response.Header.Clone()
	_, _ = response.Body.Read(body)
	for i, v := range response.Header.Clone() {
		structChan.C.Writer.Header().Add(i, v[0])
	}

	structChan.C.JSON(response.StatusCode, body)
	_ = structChan.C.BindHeader(header)
	structChan.C.Status(response.StatusCode)

	models.Log{
		BaseID:   thread.Base.Id,
		BaseName: thread.Base.Name,
		Context:  "Получен проксируемый ответ",
		InternalContext: fmt.Sprintf("Метод: %v,\n Путь: %v,\n Заголовок: %v,\n Тело: %v,\n Источник: %v,\n "+
			"Код ответа: %v\n Заголовок ответа: %v,\n Тело ответа: %v",
			request.Method, path, request.Header, request.Body, structChan.C.Request.Host, response.StatusCode,
			response.Header, response.Body),
	}.Info("Произведен проксируемый запрос")

}
