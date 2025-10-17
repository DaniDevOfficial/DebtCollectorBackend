package handlers

import (
	"dept-collector/internal/domain/dev"

	"github.com/gin-gonic/gin"
)

func RegisterDevRoutes(router *gin.RouterGroup) {
	router.GET("/helloWorld", func(c *gin.Context) {
		dev.HelloWorld(c)
	})
}
