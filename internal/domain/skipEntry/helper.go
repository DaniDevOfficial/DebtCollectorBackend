package skipEntry

import (
	"dept-collector/internal/models"
	"dept-collector/internal/responseTypes"

	"gorm.io/gorm"
)

func ApplySkipEntryFilters(db *gorm.DB, filters FilterSkipEntryRequest) *gorm.DB {
	query := db.Model(&models.SkipEntry{}).
		Joins("JOIN lessons ON lessons.id = skip_entries.lesson_id").
		Joins("JOIN classes ON classes.id = lessons.class_id").
		Joins("JOIN semesters ON semesters.id = classes.semester_id").
		Preload("User").Preload("Lesson").Preload("Amount")

	if filters.UserID != nil {
		query = query.Where("skip_entries.user_id = ?", *filters.UserID)
	}
	if filters.LessonID != nil {
		query = query.Where("skip_entries.lesson_id = ?", *filters.LessonID)
	}
	if filters.AmountID != nil {
		query = query.Where("skip_entries.amount_id = ?", *filters.AmountID)
	}
	if filters.Reason != nil {
		query = query.Where("skip_entries.reason ILIKE ?", "%"+*filters.Reason+"%")
	}
	if filters.StartDate != nil && filters.EndDate != nil {
		query = query.Where("lessons.start_date_time BETWEEN ? AND ?", *filters.StartDate, *filters.EndDate)
	} else if filters.StartDate != nil {
		query = query.Where("lessons.start_date_time >= ?", *filters.StartDate)
	} else if filters.EndDate != nil {
		query = query.Where("lessons.start_date_time <= ?", *filters.EndDate)
	}
	if filters.SemesterID != nil {
		query = query.Where("semesters.id = ?", *filters.SemesterID)
	}
	if filters.ClassID != nil {
		query = query.Where("classes.id = ?", *filters.ClassID)
	}

	return query
}

func buildSkipEntryResponse(entry *models.SkipEntry) responseTypes.SkipEntryResponse {

	return responseTypes.SkipEntryResponse{
		ID:         entry.ID,
		Reason:     entry.Reason,
		UserID:     entry.UserID,
		UserName:   entry.User.Name,
		LessonID:   entry.LessonID,
		LessonName: entry.Lesson.Name,
		Amount:     entry.Amount.Value,
		CreatedAt:  entry.CreatedAt,
		UpdatedAt:  entry.UpdatedAt,
	}
}

func buildSkipEntriesResponse(entries []models.SkipEntry) []responseTypes.SkipEntryResponse {
	responses := make([]responseTypes.SkipEntryResponse, 0, len(entries))
	for _, e := range entries {
		entry := e // avoid pointer reuse
		responses = append(responses, buildSkipEntryResponse(&entry))
	}
	return responses
}
