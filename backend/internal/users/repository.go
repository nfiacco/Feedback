package users

import (
	"feedback/internal/models"

	"gorm.io/gorm"
)

func LoadByID(db *gorm.DB, id int64) (*models.User, error) {
	var user models.User
	result := db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func Create(db *gorm.DB) (*models.User, error) {
	var user models.User
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
