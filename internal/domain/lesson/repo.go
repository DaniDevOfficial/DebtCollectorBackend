package lesson

import (
	"dept-collector/internal/models"

	"gorm.io/gorm"
)

func createNewLesson(lesson models.Lesson, db *gorm.DB) error {
	result := db.Create(&lesson)
	if result.Error != nil {
		return result.Error
	}

	result = db.Preload("Class").Where("id = ?", lesson.ID).First(&lesson)
	return result.Error
}
