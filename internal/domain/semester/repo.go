package semester

import (
	"dept-collector/internal/models"
	"gorm.io/gorm"
)

func createSemester(semester *models.Semester, db *gorm.DB) error {
	result := db.Create(semester)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}
