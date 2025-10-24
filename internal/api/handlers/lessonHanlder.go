package handlers

import (
	"dept-collector/internal/domain/lesson"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterLessonHandler(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/", func(c *gin.Context) {
		lesson.CreateNewLesson(c, db)
	})

	r.PUT("/", func(c *gin.Context) {
		lesson.EditLesson(c, db)
	})

	r.DELETE("/", func(c *gin.Context) {
		lesson.DeleteLesson(c, db)
	})

	r.GET("/", func(c *gin.Context) {
		lesson.GetSpecificLesson(c, db)
	})

	r.GET("/all", func(c *gin.Context) {
		lesson.GetFilteredLessonsWithSkipEntries(c, db)
	})

}
