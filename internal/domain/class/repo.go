package class

import (
	"dept-collector/internal/models"
	"errors"

	"github.com/google/uuid"
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

func updateClass(class *models.Class, db *gorm.DB) error {
	result := db.Save(class)
	return result.Error
}

func deleteClass(id uuid.UUID, db *gorm.DB) error {
	result := db.Delete(&models.Class{ID: id})
	return result.Error
}

func getClass(id uuid.UUID, db *gorm.DB) (models.Class, error) {
	var class models.Class
	result := db.First(&class, "id = ?", id)
	return class, result.Error
}

func getFilteredClasses(filters FilterClassRequest, db *gorm.DB) ([]models.Class, error) {
	query := db.Model(&models.Class{}).
		Joins("JOIN semesters s ON s.id = classes.semester_id")

	if filters.Name != nil {
		query = query.Where("classes.name ILIKE ?", "%"+*filters.Name+"%")
	}

	if filters.SemesterID != nil {
		query = query.Where("classes.semester_id = ?", *filters.SemesterID)
	}

	if filters.SemesterStartAfter != nil {
		query = query.Where("s.start_date >= ?", *filters.SemesterStartAfter)
	}

	if filters.SemesterEndBefore != nil {
		query = query.Where("s.end_date <= ?", *filters.SemesterEndBefore)
	}

	var classes []models.Class
	err := query.Find(&classes).Error
	return classes, err
}
