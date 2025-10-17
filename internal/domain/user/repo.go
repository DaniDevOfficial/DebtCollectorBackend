package user

import (
	"errors"

	"gorm.io/gorm"
)

func isUsernameOrEmailTaken(username string, email string, db *gorm.DB) (bool, error) {
	var user User
	result := db.Where("name = ?", username).Or("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func getUserByName(username string, db *gorm.DB) (User, error) {
	var user User
	result := db.Where("name = ?", username).First(&user)
	return user, result.Error
}
