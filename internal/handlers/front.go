package handlers

import (
	"1c_api_proxy/internal/models"
	"1c_api_proxy/internal/services/connection"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func SetDBParam(c *gin.Context) {

	var bodyData []byte
	_, err := c.Request.Body.Read(bodyData)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело запроса",
			Handler: "/db/set",
		}.Error("Не удалось прочитать тело запроса")
		c.AbortWithStatus(400)
	}

	modelBody := models.ConfSQL{}
	err = json.Unmarshal(bodyData, &modelBody)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело как JSON",
			Handler: "/db/set",
		}.Error("Не удалось прочитать тело как JSON")
		c.AbortWithStatus(400)
	}

	res := models.SetSettingsDB(&modelBody)
	if res {
		c.JSON(200, gin.H{
			"result": true,
		})
	} else {
		c.JSON(500, gin.H{
			"result":      false,
			"description": "Не удалось записать параметры",
		})
	}

}

func GetDBStatus(c *gin.Context) {

	c.JSON(200, gin.H{
		"status": models.DBConnect.IsConnected(),
	})

}

func AddInfobase(c *gin.Context) {

	var bodyData []byte
	_, err := c.Request.Body.Read(bodyData)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело запроса",
			Handler: "/infobase/add",
		}.Error("Не удалось прочитать тело запроса")
		c.AbortWithStatus(400)
	}

	modelBody := models.Infobase{}
	err = json.Unmarshal(bodyData, &modelBody)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело как JSON",
			Handler: "/infobase/add",
		}.Error("Не удалось прочитать тело как JSON")
		c.AbortWithStatus(400)
	}

	res := connection.InitService1CAPI(&modelBody)
	if res {
		c.JSON(200, gin.H{
			"result": true,
		})
	} else {
		c.JSON(500, gin.H{
			"result":      false,
			"description": "Не удалось создать, возможно отсутствует имя или URL",
		})
	}

}

func EditInfobase(c *gin.Context) {

	var bodyData []byte
	_, err := c.Request.Body.Read(bodyData)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело запроса",
			Handler: "/infobase/add",
		}.Error("Не удалось прочитать тело запроса")
		c.AbortWithStatus(400)
	}

	modelBody := models.Infobase{}
	err = json.Unmarshal(bodyData, &modelBody)
	if err != nil {
		models.Log{
			Context: err.Error(),
			Comment: "Не удалось прочитать тело как JSON",
			Handler: "/infobase/add",
		}.Error("Не удалось прочитать тело как JSON")
		c.AbortWithStatus(400)
	}

	res := models.Connections.EditThread(&modelBody)

	if res {
		c.JSON(200, gin.H{
			"result": true,
		})
	} else {
		c.JSON(500, gin.H{
			"result":      false,
			"description": "Не удалось найти или заменить",
		})
	}

}

func DeleteInfobase(c *gin.Context) {

	nameInfobase := c.Request.Header.Get("infobase")

	if nameInfobase == "" {
		models.Log{
			Comment: "Не заполнено имя информационной базы с заголовке",
			Handler: "/infobase/status",
		}.Error("Не заполнено имя информационной базы с заголовке")
		c.AbortWithStatus(400)
	}

	findInfobaseThread := models.Connections.FindThreadConnectByName(nameInfobase)
	if findInfobaseThread == nil {
		c.JSON(400, gin.H{
			"result":      false,
			"description": "Информационная база не была найдена",
		})
	}

	res := models.Connections.DeleteThread(findInfobaseThread)
	if res {
		c.JSON(200, gin.H{
			"result": true,
		})
	} else {
		c.JSON(500, gin.H{
			"result":      false,
			"description": "Не удалось создать, возможно отсутствует имя или URL",
		})
	}

}

func GetInfobases(c *gin.Context) {

	list := models.Connections.GetInfobasesList()

	body, err := json.Marshal(list)
	if err != nil {
		c.JSON(500, gin.H{
			"result":      false,
			"description": "Не удалось сформировать ответ",
		})
	}

	c.JSON(200, body)

}

func ReloadConnect(c *gin.Context) {

	nameInfobase := c.Request.Header.Get("infobase")

	if nameInfobase == "" {
		models.Log{
			Comment: "Не заполнено имя информационной базы с заголовке",
			Handler: "/infobase/status",
		}.Error("Не заполнено имя информационной базы с заголовке")
		c.AbortWithStatus(400)
	}

	findInfobaseThread := models.Connections.FindThreadConnectByName(nameInfobase)
	if findInfobaseThread == nil {
		c.JSON(400, gin.H{
			"result":      false,
			"description": "Информационная база не была найдена",
		})
	}

	if !findInfobaseThread.ChanIsClosed() {
		findInfobaseThread.CloseLoop()
	}
	connection.RestartLoop(findInfobaseThread)
	c.JSON(200, gin.H{
		"result": true,
	})
}

func GetStatusInfobase(c *gin.Context) {

	nameInfobase := c.Request.Header.Get("infobase")

	if nameInfobase == "" {
		models.Log{
			Comment: "Не заполнено имя информационной базы с заголовке",
			Handler: "/infobase/status",
		}.Error("Не заполнено имя информационной базы с заголовке")
		c.AbortWithStatus(400)
	}

	findInfobaseThread := models.Connections.FindThreadConnectByName(nameInfobase)
	if findInfobaseThread == nil {
		c.JSON(400, gin.H{
			"result":      false,
			"description": "Информационная база не была найдена",
		})
	}

	c.JSON(200, gin.H{
		"result": findInfobaseThread.ChanIsClosed(),
	})

}

func GetLogs(c *gin.Context) {

}
