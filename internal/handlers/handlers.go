package handlers

import (
	"1c_api_proxy/internal/models"
	connection2 "1c_api_proxy/internal/services/connection"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Proxy(c *gin.Context) {

	infobaseName := c.GetHeader("infobase")
	connection := models.Connections.FindThreadConnectByName(infobaseName)

	if connection == nil {
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
		"message": "Сервис прокси для информационных баз 1С. Релиз 0.0.1 (альфа) \n " +
			"Используйте /proxy/: чтобы сделать запрос \n обязательно в заголовок укажите параметр" +
			" 'infobase' с именем информационной базы",
	})

}
