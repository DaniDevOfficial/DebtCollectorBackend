package lesson

import (
	"dept-collector/internal/models"

	"github.com/google/uuid"
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

func updateLesson(lesson models.Lesson, db *gorm.DB) error {
	result := db.Save(&lesson)
	if result.Error != nil {
		return result.Error
	}
	result = db.Preload("Class").Where("id = ?", lesson.ID).First(&lesson)
	return result.Error
}

func deleteLesson(id uuid.UUID, db *gorm.DB) error {
	result := db.Delete(&models.Lesson{ID: id})
	return result.Error
}

func getLesson(id uuid.UUID, db *gorm.DB) (models.Lesson, error) {
	var lesson models.Lesson
	result := db.Preload("Class.Semester").Where("id = ?", id).First(&lesson)
	return lesson, result.Error
}

func getAllLessonsWithSkipEntries(filters FilterLessonRequest, db *gorm.DB) ([]models.Lesson, error) {
	var lessons []models.Lesson
	query := ApplyLessonFilters(filters, db)
	err := query.
		Preload("Class.Semester").
		Preload("SkipEntries.User").
		Preload("SkipEntries.Amount").
		Find(&lessons).
		Error
	return lessons, err
}
