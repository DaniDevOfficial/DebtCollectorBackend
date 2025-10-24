package lesson

import (
	"dept-collector/internal/models"

	"gorm.io/gorm"
)

func ApplyLessonFilters(filters FilterLessonRequest, db *gorm.DB) *gorm.DB {
	query := db.Model(&models.Lesson{}).
		Joins("LEFT JOIN skip_entry se ON lesson.id = se.lesson_id").
		Joins("JOIN user u ON u.id = se.user_id").
		Preload("user").Preload("skip_entry")

	if filters.ClassID != nil {
		query = query.Where("lesson.class_id = ?", filters.ClassID)
	}

	if filters.StartDate != nil {
		query = query.Where("lesson.start_date >= ?", filters.StartDate)
	}
	if filters.EndDate != nil {
		query = query.Where("lesson.end_date <= ?", filters.EndDate)
	}

	if filters.Name != nil {
		query = query.Where("lesson.name ILIKE ?", "%"+*filters.Name+"%")
	}

	return query

}
