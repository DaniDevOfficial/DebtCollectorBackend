package api

import (
	"dept-collector/internal/api/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	basePath := router.Group("/api")

	handlers.RegisterDevRoutes(basePath.Group("/dev"))
	handlers.RegisterUserRoutes(basePath.Group("/user"), db)
	handlers.RegisterSkipEntryRoutes(basePath.Group("/skips"), db)

	return router
}
