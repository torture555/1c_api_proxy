package handlers

import (
	"1c_api_proxy/internal/models"
	"1c_api_proxy/internal/services/connection"
	"1c_api_proxy/internal/transport/rest/front"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
)

func SetDBParam(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело запроса",
			Handler: front.Database + "/" + front.SetDBParams,
		}.Error("Не удалось прочитать тело запроса")
		c.AbortWithStatus(400)
		c.Next()
		return
	}

	modelBody := models.ConfSQL{}
	err = json.Unmarshal(body, &modelBody)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело как JSON",
			Handler: front.Database + "/" + front.SetDBParams,
		}.Error("Не удалось прочитать тело как JSON")
		c.AbortWithStatus(400)
		c.Next()
		return
	}

	res := models.SetSettingsDB(&modelBody)
	if res {
		c.JSON(200, gin.H{
			"result": true,
		})
		c.Next()
		return
	} else {
		c.JSON(500, gin.H{
			"result":      false,
			"description": "Не удалось записать параметры",
		})
		c.Next()
		return
	}

}

func GetDBParam(c *gin.Context) {

	list, err := models.GetSettingsDB()
	if err != nil {
		c.JSON(500, gin.H{
			"result":      false,
			"description": "Не удалось сформировать ответ",
		})
		c.Next()
		return
	}

	c.JSON(200, gin.H{
		"host":   list.Host,
		"port":   list.Port,
		"login":  list.Login,
		"DBName": list.DBName,
	})
	c.Next()
	return

}

func GetDBStatus(c *gin.Context) {

	c.JSON(200, gin.H{
		"status": models.DBConnect.IsConnected(),
	})
	c.Next()
	return

}

func AddInfobase(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело запроса",
			Handler: front.Database + "/" + front.SetDBParams,
		}.Error("Не удалось прочитать тело запроса")
		c.AbortWithStatus(400)
		c.Next()
		return
	}

	modelBody := models.Infobase{}
	err = json.Unmarshal(body, &modelBody)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело как JSON",
			Handler: front.Infobases + "/" + front.AddInfobase,
		}.Error("Не удалось прочитать тело как JSON")
		c.AbortWithStatus(400)
		c.Next()
		return
	}

	res := connection.InitService1CAPI(&modelBody)
	if res {
		c.JSON(200, gin.H{
			"result": true,
		})
		c.Next()
		return
	} else {
		c.JSON(500, gin.H{
			"result":      false,
			"description": "Не удалось создать, возможно отсутствует имя или URL",
		})
		c.Next()
		return
	}

}

func EditInfobase(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело запроса",
			Handler: front.Database + "/" + front.SetDBParams,
		}.Error("Не удалось прочитать тело запроса")
		c.AbortWithStatus(400)
		c.Next()
		return
	}

	modelBody := models.Infobase{}
	err = json.Unmarshal(body, &modelBody)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело как JSON",
			Handler: front.Infobases + "/" + front.EditInfobase,
		}.Error("Не удалось прочитать тело как JSON")
		c.AbortWithStatus(400)
		c.Next()
		return
	}

	res := models.Connections.EditThread(&modelBody)

	if res {
		c.JSON(200, gin.H{
			"result": true,
		})
		c.Next()
		return
	} else {
		c.JSON(500, gin.H{
			"result":      false,
			"description": "Не удалось найти или заменить",
		})
		c.Next()
		return
	}

}

func DeleteInfobase(c *gin.Context) {

	nameInfobase := c.Request.Header.Get("Infobase")

	if nameInfobase == "" {
		models.Log{
			Comment: "Не заполнено имя информационной базы с заголовке",
			Handler: front.Infobases + "/" + front.DeleteInfobase,
		}.Error("Не заполнено имя информационной базы с заголовке")
		c.AbortWithStatus(400)
		c.Next()
		return
	}

	findInfobaseThread := models.Connections.FindThreadConnectByName(nameInfobase)
	if findInfobaseThread == nil {
		c.JSON(400, gin.H{
			"result":      false,
			"description": "Информационная база не была найдена",
		})
		c.Next()
		return
	}

	res := models.Connections.DeleteThread(findInfobaseThread)
	if res {
		c.JSON(200, gin.H{
			"result": true,
		})
		c.Next()
		return
	} else {
		c.JSON(500, gin.H{
			"result":      false,
			"description": "Не удалось создать, возможно отсутствует имя или URL",
		})
		c.Next()
		return
	}

}

func GetInfobases(c *gin.Context) {

	list := models.Connections.GetInfobasesList()
	c.JSON(200, list.Bases)
	c.Next()
	return

}

func ReloadConnect(c *gin.Context) {

	nameInfobase := c.Request.Header.Get("Infobase")

	if nameInfobase == "" {
		models.Log{
			Comment: "Не заполнено имя информационной базы с заголовке",
			Handler: front.Infobases + front.ReloadConnect,
		}.Error("Не заполнено имя информационной базы с заголовке")
		c.AbortWithStatus(400)
		c.Next()
		return
	}

	findInfobaseThread := models.Connections.FindThreadConnectByName(nameInfobase)
	if findInfobaseThread == nil {
		c.JSON(400, gin.H{
			"result":      false,
			"description": "Информационная база не была найдена",
		})
		c.Next()
		return
	}

	connection.RestartLoop(findInfobaseThread)
	c.JSON(200, gin.H{
		"result": true,
	})
	c.Next()
	return
}

func GetStatusInfobase(c *gin.Context) {

	nameInfobase := c.Request.Header.Get("Infobase")

	if nameInfobase == "" {
		models.Log{
			Comment: "Не заполнено имя информационной базы с заголовке",
			Handler: front.Infobases + "/" + front.Status,
		}.Error("Не заполнено имя информационной базы с заголовке")
		c.AbortWithStatus(400)
		c.Next()
		return
	}

	findInfobaseThread := models.Connections.FindThreadConnectByName(nameInfobase)
	if findInfobaseThread == nil {
		c.JSON(400, gin.H{
			"result":      false,
			"description": "Информационная база не была найдена",
		})
		c.Next()
		return
	}

	c.JSON(200, gin.H{
		"result": !findInfobaseThread.ChanIsClosed(),
	})
	c.Next()
	return

}

func GetLogs(c *gin.Context) {

	logs := models.GetLogs()
	c.JSON(200, logs)

}
