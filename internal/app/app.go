package app

import (
	"1c_api_proxy/internal/handlers"
	"1c_api_proxy/internal/middleware"
	api_v1 "1c_api_proxy/internal/transport/rest/v1"
	timeout "github.com/cyj19/gin-timeout"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
	"time"
)

func StartRouteProxy(wg *sync.WaitGroup) {

	engine := gin.Default()

	timeoutTime := 3 * time.Second
	opt := timeout.Option{
		Timeout: &timeoutTime,
		Code:    408,
		Msg:     "Timeout abort",
	}

	engine.Use(middleware.Cors)
	engine.Use(timeout.ContextTimeout(opt))

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

	if err != nil {
		wg.Done()
	}

}

func initGetPostProxy(group *gin.RouterGroup) {
	group.GET("/", handlers.Proxy)
	group.POST("/", handlers.Proxy)
}
