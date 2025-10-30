package handlers

import (
	"dept-collector/internal/domain/class"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterClassRoutes(router *gin.RouterGroup, db *gorm.DB) {

	router.POST("/create", func(c *gin.Context) {
		class.CreateNewClass(c, db)
	})
}
