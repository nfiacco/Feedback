package external_profiles

import (
	"feedback/internal/models"

	"gorm.io/gorm"
)

func LoadByExternalID(db *gorm.DB, externalID int64) (*models.ExternalProfile, error) {
	var externalProfile models.ExternalProfile
	result := db.First(&externalProfile)
	if result.Error != nil {
		return nil, result.Error
	}

	return &externalProfile, nil
}

func Create(db *gorm.DB, externalID string, userID int64) (*models.ExternalProfile, error) {
	externalProfile := models.ExternalProfile{
		ExternalID: externalID,
		UserID:     userID,
	}
	result := db.Create(&externalProfile)
	if result.Error != nil {
		return nil, result.Error
	}

	return &externalProfile, nil
}
