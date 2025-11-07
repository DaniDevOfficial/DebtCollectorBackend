package lesson

import (
	"dept-collector/internal/models"
	"dept-collector/internal/responseTypes"

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
