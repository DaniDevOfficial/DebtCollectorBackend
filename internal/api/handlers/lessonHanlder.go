package handlers

import (
	"dept-collector/internal/domain/lesson"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterLessonHandler(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/lesson/create", func(c *gin.Context) {
		lesson.CreateNewLesson(c, db)
	})
}
