package user_identities

import (
	"feedback/internal/models"

	"gorm.io/gorm"
)

func Create(db *gorm.DB, firstName string, lastName string, userID int64) (*models.UserIdentity, error) {
	userIdentity := models.UserIdentity{
		FirstName: firstName,
		LastName:  lastName,
		UserID:    userID,
	}
	result := db.Create(&userIdentity)
	if result.Error != nil {
		return nil, result.Error
	}

	return &userIdentity, nil
}
