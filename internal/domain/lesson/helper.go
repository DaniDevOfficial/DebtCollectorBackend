package lesson

import (
	"dept-collector/internal/models"
	"dept-collector/internal/responseTypes"

	"gorm.io/gorm"
)

func ApplyLessonFilters(filters FilterLessonRequest, db *gorm.DB) *gorm.DB {
	query := db.Model(&models.Lesson{}).
		// alias the lessons table as "l"
		Table("lessons AS l").
		Joins("LEFT JOIN skip_entries se ON l.id = se.lesson_id").
		Joins("LEFT JOIN users u ON u.id = se.user_id").
		Preload("SkipEntries").
		Preload("Class").
		Preload("Class.Semester")

	if filters.ClassID != nil {
		query = query.Where("l.class_id = ?", filters.ClassID)
	}

	if filters.StartDate != nil {
		query = query.Where("l.start_date_time >= ?", filters.StartDate)
	}

	if filters.EndDate != nil {
		query = query.Where("l.end_date_time <= ?", filters.EndDate)
	}

	if filters.Name != nil {
		query = query.Where("l.name ILIKE ?", "%"+*filters.Name+"%")
	}

	return query
}

func buildFilteredLessonsResponse(lessons []models.Lesson) []responseTypes.FilteredLesson {
	filtered := make([]responseTypes.FilteredLesson, 0, len(lessons))

	for _, lesson := range lessons {
		skipResponses := make([]responseTypes.SkipEntryResponse, 0, len(lesson.SkipEntries))
		for _, skipEntry := range lesson.SkipEntries {

			skipResponses = append(skipResponses, responseTypes.SkipEntryResponse{
				ID:         skipEntry.ID,
				Reason:     skipEntry.Reason,
				UserID:     skipEntry.UserID,
				UserName:   skipEntry.User.Name,
				LessonID:   skipEntry.LessonID,
				LessonName: lesson.Name,
				Amount:     skipEntry.Amount.Value,
				CreatedAt:  skipEntry.CreatedAt,
				UpdatedAt:  skipEntry.UpdatedAt,
			})
		}

		filtered = append(filtered, responseTypes.FilteredLesson{
			ID:            lesson.ID,
			Name:          lesson.Name,
			StartDateTime: lesson.StartDateTime,
			EndDateTime:   lesson.EndDateTime,
			ClassID:       lesson.Class.ID,
			ClassName:     lesson.Class.Name,
			SemesterID:    lesson.Class.Semester.ID,
			SemesterName:  lesson.Class.Semester.Name,
			SkipEntries:   skipResponses,
			CreatedAt:     lesson.CreatedAt,
			UpdatedAt:     lesson.UpdatedAt,
		})
	}

	return filtered
}
