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
	router.POST("", func(c *gin.Context) {
		skipEntry.CreateNewSkipEntry(c, db)
	})

	router.PUT("", func(c *gin.Context) {
		skipEntry.EditSkipEntry(c, db)
	})

	router.DELETE("", func(c *gin.Context) {
		skipEntry.DeleteSkipEntry(c, db)
	})

	router.GET("", func(c *gin.Context) {
		skipEntry.GetSpecificSkipEntry(c, db)
	})

	router.GET("/filtered", func(c *gin.Context) {
		skipEntry.GetFilteredSkipEntries(c, db)
	})
}
