package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidProxy(c *gin.Context) {

	if c.GetHeader("infobase") == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

}
