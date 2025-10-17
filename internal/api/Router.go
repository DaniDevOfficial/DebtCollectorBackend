package api

import (
	"dept-collector/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	basePath := router.Group("/api")

	handlers.RegisterDevRoutes(basePath.Group("/dev"))

	return router
}
