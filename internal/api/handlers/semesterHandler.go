package handlers

import (
	"dept-collector/internal/domain/semester"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterSemesterRoutes(router *gin.RouterGroup, db *gorm.DB) {

	router.POST("/create", func(c *gin.Context) {
		semester.CreateNewSemester(c, db)
	})
}
