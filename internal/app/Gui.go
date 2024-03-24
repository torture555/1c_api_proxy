package app

import (
	"1c_api_proxy/internal/handlers"
	"1c_api_proxy/internal/middleware"
	"1c_api_proxy/internal/transport/rest/front"
	timeout "github.com/cyj19/gin-timeout"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
	"time"
)

func StartBackend(wg *sync.WaitGroup) {

	engine := gin.Default()

	timeoutTime := 100 * time.Millisecond
	opt := timeout.Option{
		Timeout: &timeoutTime,
		Code:    408,
		Msg:     "Timeout abort",
	}

	engine.Use(middleware.Cors)
	engine.Use(timeout.ContextTimeout(opt))

	engine.GET("/"+front.Infobases, handlers.GetInfobases)
	engine.POST("/"+front.Infobases+"/"+front.AddInfobase, handlers.AddInfobase)
	engine.POST("/"+front.Infobases+"/"+front.EditInfobase, handlers.EditInfobase)
	engine.GET("/"+front.Infobases+"/"+front.DeleteInfobase, handlers.DeleteInfobase)
	engine.GET("/"+front.Infobases+"/"+front.Status, handlers.GetStatusInfobase)
	engine.GET("/"+front.Infobases+"/"+front.ReloadConnect, handlers.ReloadConnect)

	engine.GET("/"+front.Database+"/"+front.StatusDB, handlers.GetDBStatus)
	engine.GET("/"+front.Database+"/"+front.GetDBParams, handlers.GetDBParam)
	engine.POST("/"+front.Database+"/"+front.SetDBParams, handlers.SetDBParam)

	engine.GET("/"+front.Logs+"/"+front.GetLog, handlers.GetLogs)

	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"ok": "ok",
		})
	})

	err := engine.Run(":" + strconv.Itoa(10001))

	if err != nil {
		wg.Done()
	}

}

func StartFrontEnd(wg *sync.WaitGroup) {

	engine := gin.Default()

	engine.LoadHTMLFiles("dist/index.html")
	engine.StaticFile("/favicon.ico", "./dist/favicon.ico")
	engine.Static("/assets", "./dist/assets")
	engine.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{
			"title": "Панель управления прокси сервером 1С",
		})
	})

	err := engine.Run(":" + strconv.Itoa(11021))

	if err != nil {
		wg.Done()
	}

}
