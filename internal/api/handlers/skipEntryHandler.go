package handlers

import (
	"dept-collector/internal/domain/skipEntry"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterSkipEntryRoutes(router *gin.RouterGroup, db *gorm.DB) {
	registerCRUDSkipEntryRoutes(router, db)
}

func registerCRUDSkipEntryRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/create", func(c *gin.Context) {
		skipEntry.CreateNewSkipEntry(c, db)
	})
}
