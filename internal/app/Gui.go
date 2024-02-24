package app

import (
	"1c_api_proxy/internal/handlers"
	"1c_api_proxy/internal/transport/rest/front"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
)

func StartBackend(wg *sync.WaitGroup) {

	engine := gin.Default()

	engine.GET("/"+front.Infobases, handlers.GetInfobases)
	engine.POST("/"+front.Infobases+"/"+front.AddInfobase, handlers.AddInfobase)
	engine.POST("/"+front.Infobases+"/"+front.EditInfobase, handlers.EditInfobase)
	engine.GET("/"+front.Infobases+"/"+front.DeleteInfobase, handlers.DeleteInfobase)
	engine.GET("/"+front.Infobases+"/"+front.Status, handlers.GetStatusInfobase)
	engine.GET("/"+front.Infobases+"/"+front.ReloadConnect, handlers.ReloadConnect)

	engine.GET("/"+front.Database+"/"+front.StatusDB, handlers.GetDBStatus)
	engine.POST("/"+front.Database+"/"+front.SetDBParams, handlers.SetDBParam)

	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"ok": "ok",
		})
	})

	err := engine.Run(":" + strconv.Itoa(11021))

	if err != nil {
		wg.Done()
	}

}

func StartFrontEnd(wg *sync.WaitGroup) {

	engine := gin.Default()

	engine.LoadHTMLFiles("dist/index.html")

	engine.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{})
	})

	err := engine.Run(":" + strconv.Itoa(11021))

	if err != nil {
		wg.Done()
	}

}
