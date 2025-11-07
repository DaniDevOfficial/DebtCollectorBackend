package handlers

import (
	"dept-collector/internal/domain/semester"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterSemesterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	router.POST("/", func(c *gin.Context) {
		semester.CreateNewSemester(c, db)
	})

	router.PUT("/", func(c *gin.Context) {
		semester.EditSemester(c, db)
	})
	router.GET("/", func(c *gin.Context) {
		semester.GetSpecificSemester(c, db)
	})
	router.DELETE("/", func(c *gin.Context) {
		semester.DeleteSemester(c, db)
	})
}
