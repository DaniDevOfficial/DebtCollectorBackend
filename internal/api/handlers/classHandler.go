package handlers

import (
	"dept-collector/internal/domain/class"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterClassRoutes(router *gin.RouterGroup, db *gorm.DB) {

	router.POST("", func(c *gin.Context) {
		class.CreateNewClass(c, db)
	})

	router.PUT("", func(c *gin.Context) {
		class.EditClass(c, db)
	})

	router.GET("", func(c *gin.Context) {
		class.GetClass(c, db)
	})

	router.DELETE("", func(c *gin.Context) {
		class.DeleteClass(c, db)
	})

	router.GET("/filtered", func(c *gin.Context) {
		class.GetFilteredClasses(c, db)
	})

}
