package user

import (
	"dept-collector/internal/models"
	"errors"

	"gorm.io/gorm"
)

func isUsernameOrEmailTaken(username string, email string, db *gorm.DB) (bool, error) {
	var user models.User
	result := db.Where("name = ?", username).Or("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func getUserByName(username string, db *gorm.DB) (models.User, error) {
	var user models.User
	result := db.Where("name = ?", username).First(&user)
	return user, result.Error
}

func createNewUser(newUser *models.User, db *gorm.DB) error {
	result := db.Create(newUser)
	return result.Error
}
