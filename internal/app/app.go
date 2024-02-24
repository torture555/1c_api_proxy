package app

import (
	"1c_api_proxy/internal/handlers"
	api_v1 "1c_api_proxy/internal/transport/rest/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func StartRouteProxy(wg *sync.WaitGroup) {

	engine := gin.Default()

	s := &http.Server{
		Addr:              ":" + strconv.Itoa(10000),
		Handler:           engine,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	groupProxy := engine.Group(api_v1.PathProxy_Proxy)
	initGetPostProxy(groupProxy)

	groupFirstLevel := groupProxy.Group("/:first")
	initGetPostProxy(groupFirstLevel)

	groupSecondLevel := groupFirstLevel.Group("/:second")
	initGetPostProxy(groupSecondLevel)

	groupThirdLevel := groupSecondLevel.Group("/:third")
	initGetPostProxy(groupThirdLevel)

	groupFourthLevel := groupThirdLevel.Group("/:fourth")
	initGetPostProxy(groupFourthLevel)

	groupFifthLevel := groupFourthLevel.Group("/:fifth")
	initGetPostProxy(groupFifthLevel)

	groupSixthLevel := groupFifthLevel.Group("/:sixth")
	initGetPostProxy(groupSixthLevel)

	groupSeventhLevel := groupSixthLevel.Group("/:seventh")
	initGetPostProxy(groupSeventhLevel)

	engine.GET("/", handlers.Help)

	err := engine.Run(":" + strconv.Itoa(10000))

	err = s.ListenAndServe()

	if err != nil {
		wg.Done()
	}

}

func initGetPostProxy(group *gin.RouterGroup) {
	group.GET("/", handlers.Proxy)
	group.POST("/", handlers.Proxy)
}
