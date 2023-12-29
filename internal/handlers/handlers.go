package handlers

import (
	"1c_api_proxy/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Proxy(ctx *gin.Context) {

	c := ctx.Copy()

	infobaseName := c.GetHeader("infobase")
	connection := models.Connections.FindThreadConnectByName(infobaseName)

	if connection == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"description": "Информационная база не найдена",
		})
	}

	chanConnect := models.ModelChanConnect{
		Request: *c.Request,
	}

	connection.ChanResponseRequest <- chanConnect
	result := <-connection.ChanResponseRequest

	if !result.Result {
		c.JSON(http.StatusInternalServerError, gin.H{
			"description": "Не удалось получить ответ от инф.базы",
		})
	}

	_ = c.BindHeader(result.Response.Header)
	_ = c.Bind(result.Response.Body)
	c.Status(result.Response.StatusCode)

}

func Help(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
		"message": "Сервис прокси для информационных баз 1С. Релиз 0.0.1 (альфа) \n " +
			"Используйте /proxy/: чтобы сделать запрос \n обязательно в заголовок укажите параметр" +
			" 'infobase' с именем информационной базы",
	})

}
