package skipEntry

import (
	"dept-collector/internal/models"

	"github.com/google/uuid"
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

func deleteSkipEntry(id uuid.UUID, db *gorm.DB) error {
	result := db.Delete(&models.SkipEntry{ID: id})
	return result.Error
}

func getSpecificEntry(id uuid.UUID, db *gorm.DB) (*models.SkipEntry, error) {
	var entry models.SkipEntry
	result := db.
		Preload("User").
		Preload("Lesson").
		Preload("Amount").
		First(&entry, "id = ?", id)

	return &entry, result.Error
}

func getAllEntries(filters FilterSkipEntryRequest, db *gorm.DB) ([]models.SkipEntry, error) {
	var entries []models.SkipEntry
	query := ApplySkipEntryFilters(db, filters)
	err := query.Find(&entries).Error
	return entries, err
}
