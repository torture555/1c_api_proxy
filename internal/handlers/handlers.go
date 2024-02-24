package handlers

import (
	"1c_api_proxy/internal/models"
	connection2 "1c_api_proxy/internal/services/connection"
	api_v1 "1c_api_proxy/internal/transport/rest/v1"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Proxy(c *gin.Context) {

	baseName := strings.ReplaceAll(c.Request.RequestURI, api_v1.PathProxy_Proxy, "")
	baseNameArr := strings.Split(baseName, "/")
	infobaseName := ""
	for _, v := range baseNameArr {
		if v != "" {
			infobaseName = v
			break
		}
	}
	if infobaseName == "" {
		models.Log{
			BaseName: "?",
			Context:  fmt.Sprintf("URL: %v", c.Request.URL),
			Comment:  "Не найдена инф.база",
			Handler:  "/proxy",
		}.Info("")
		c.JSON(http.StatusNotFound, gin.H{
			"description": "Информационная база не найдена",
		})
	}

	connection := models.Connections.FindThreadConnectByName(infobaseName)

	if connection == nil {
		models.Log{
			BaseName: "?",
			Context:  fmt.Sprintf("URL: %v", c.Request.URL),
			Comment:  "Не найдена инф.база",
			Handler:  "/proxy",
		}.Info("")
		c.JSON(http.StatusNotFound, gin.H{
			"description": "Информационная база не найдена",
		})
	}

	chanConnect := models.ModelChanConnect{C: c}

	connection2.ProxyRequest(connection, &chanConnect)

}

func Help(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
		"message": "Сервис прокси для информационных баз 1С. Релиз 0.1.0 (Альфа) \n " +
			"Используйте /proxy/<Имя базы>/... : чтобы сделать запрос, укажите имя инфомационной базы \n" +
			"Для администрирования обратитесь через браузер по порту фронт приложения (по умолчанию порт 11021)",
	})

}
