package class

import (
	"dept-collector/internal/models"
	"errors"
	"gorm.io/gorm"
)

var ErrSemesterNotFound = errors.New("semester not found")

func createClass(class *models.Class, db *gorm.DB) error {
	var semester models.Semester
	if err := db.First(&semester, "id = ?", class.SemesterID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrSemesterNotFound
		}
		return err
	}

	result := db.Create(class)
	return result.Error
}
