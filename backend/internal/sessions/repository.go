package sessions

import (
	"feedback/internal/models"

	"gorm.io/gorm"
)

func LoadByToken(db *gorm.DB, token string) (*models.Session, error) {
	var session models.Session
	result := db.First(&session, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}

	return &session, nil
}
