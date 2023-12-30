package app

import (
	"1c_api_proxy/internal/handlers"
	"1c_api_proxy/internal/middleware"
	"1c_api_proxy/internal/models"
	api_v1 "1c_api_proxy/internal/transport/rest/v1"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"os"
	"strconv"
	"time"
)

func StartRouteProxy() {

	fileConf := "config/app.json"

	configModel := models.ConfApp{}
	dataFile, err := os.ReadFile(fileConf)
	if err != nil {
		var raw models.Log
		raw.Context = err.Error()
		raw.Comment = "Чтение конфига приложения"
		raw.Error("Не удалось прочитать или найти app.json")
		panic(raw)
	}
	err = json.Unmarshal(dataFile, &configModel)
	if err != nil {
		var raw models.Log
		raw.Context = err.Error()
		raw.Comment = "Чтение конфига приложения"
		raw.Error("Не удалось прочитать или найти app.json")
		panic(raw)
	}

	engine := gin.Default()

	s := &http.Server{
		Addr:              ":" + strconv.Itoa(configModel.Port),
		Handler:           engine,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	groupProxy := engine.Group(api_v1.PathProxy_Proxy)
	groupProxy.Use(middleware.ValidProxy)
	initGetPostProxy(groupProxy)

	groupFirstLevel := groupProxy.Group("/:first")
	initGetPostProxy(groupFirstLevel)

	groupSecondLevel := groupFirstLevel.Group("/:second")
	initGetPostProxy(groupSecondLevel)

	groupThirdLevel := groupSecondLevel.Group("/:third")
	initGetPostProxy(groupThirdLevel)

	groupFourthLevel := groupThirdLevel.Group("/:fourth")
	initGetPostProxy(groupFourthLevel)

	engine.GET("/", handlers.Help)

	_ = engine.Run(":" + strconv.Itoa(configModel.Port))

	_ = s.ListenAndServe()

}

func initGetPostProxy(group *gin.RouterGroup) {
	group.GET("/", handlers.Proxy)
	group.POST("/", handlers.Proxy)
}
