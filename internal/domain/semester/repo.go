package semester

import (
	"dept-collector/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func createSemester(semester *models.Semester, db *gorm.DB) error {
	result := db.Create(semester)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

func updateSemester(semester *models.Semester, db *gorm.DB) error {
	result := db.Save(semester)
	return result.Error
}

func getSemester(id uuid.UUID, db *gorm.DB) (models.Semester, error) {
	var semester models.Semester
	result := db.First(&semester, "id = ?", id)
	return semester, result.Error
}

func deleteSemester(id uuid.UUID, db *gorm.DB) error {
	result := db.Delete(&models.Semester{}, id)
	return result.Error
}
