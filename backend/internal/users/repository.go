package users

import (
	"errors"
	"feedback/internal/emails"
	"feedback/internal/external_profiles"
	"feedback/internal/models"
	"feedback/internal/user_identities"
	"fmt"
	"math/big"
	"math/rand"

	"gorm.io/gorm"
)

// Maximum of 62^8 guarantees number will be at most 8 digits in base
const MAX_RANDOM = 218340105584896

type ExternalUserInfo struct {
	ExternalID string
	Email      string
	FirstName  string
	LastName   string
}

type UserAndIdentity struct {
	*models.User
	*models.UserIdentity
}

// Generates a mostly-unique 8 character string. Must check DB since collisions are possible
// There is a race condition here, since the unique key can be used by another concurrent
// process. In this case, it is ok to error since this is extremely rare.
func generateUniqueKey(db *gorm.DB) (*string, error) {
	for {
		unique := rand.Int63n(MAX_RANDOM)
		uniqueStr := big.NewInt(unique).Text(62)
		uniqueKey := fmt.Sprintf("%08s", uniqueStr)

		_, err := LoadByFeedbackKey(db, uniqueKey)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &uniqueKey, nil
			} else {
				return nil, err
			}
		}
	}
}

func LoadByExternalID(db *gorm.DB, externalID string) (*models.User, error) {
	var user models.User
	result := db.Table("users").
		Joins("JOIN external_profiles ON external_profiles.user_id = users.id").
		Where("external_profiles.external_id = ?", externalID).
		Take(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func LoadUserAndIdentityByID(db *gorm.DB, userID int64) (*UserAndIdentity, error) {
	var userAndIdentity UserAndIdentity
	result := db.Table("users").
		Joins("JOIN user_identities ON user_identities.user_id = users.id").
		Where("users.id = ?", userID).
		Take(&userAndIdentity)
	if result.Error != nil {
		return nil, result.Error
	}

	return &userAndIdentity, nil
}

func LoadByFeedbackKey(db *gorm.DB, feedbackKey string) (*models.User, error) {
	var user models.User
	result := db.Take(&user, "feedback_key = ?", feedbackKey)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func create(db *gorm.DB) (*models.User, error) {
	uniqueKey, err := generateUniqueKey(db)
	if err != nil {
		return nil, err
	}

	user := models.User{
		FeedbackKey: *uniqueKey,
	}

	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func CreateUser(db *gorm.DB, externalUserInfo *ExternalUserInfo) (*models.User, error) {
	user, err := create(db)
	if err != nil {
		return nil, err
	}

	_, err = external_profiles.Create(db, externalUserInfo.ExternalID, user.ID)
	if err != nil {
		return nil, err
	}

	_, err = emails.Create(db, externalUserInfo.Email, user.ID)
	if err != nil {
		return nil, err
	}

	_, err = user_identities.Create(db, externalUserInfo.FirstName, externalUserInfo.LastName, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetOrCreateUser(db *gorm.DB, externalUserInfo *ExternalUserInfo) (*models.User, error) {
	existingUser, err := LoadByExternalID(db, externalUserInfo.ExternalID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err == nil {
		return existingUser, nil
	}

	user, err := CreateUser(db, externalUserInfo)
	if err != nil {
		return nil, err
	}

	return user, nil
}
