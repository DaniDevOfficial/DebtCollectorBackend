package skipEntry

import (
	"dept-collector/internal/models"

	"gorm.io/gorm"
)

func createSkipEntry(entry models.SkipEntry, db *gorm.DB) error {
	result := db.Create(entry)
	return result.Error
}

func updateSkipEntry(entry *models.SkipEntry, db *gorm.DB) error {
	result := db.Save(entry)
	return result.Error
}
