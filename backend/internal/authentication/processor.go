package authentication

import (
	"feedback/internal/errors"
	"feedback/internal/models"
	"feedback/internal/sessions"
	"net/http"

	"gorm.io/gorm"
)

func Authenticate(db *gorm.DB, r *http.Request) (*models.Session, error) {
	token := r.Header.Get("X-Session-Token")
	if token == "" {
		return nil, gorm.ErrRecordNotFound
	}

	session, err := sessions.LoadValidByToken(db, token)
	if err != nil {
		return nil, err
	}

	refreshed, err := sessions.Refresh(db, session)
	if err != nil {
		return nil, err
	}

	return refreshed, nil
}

func IsAuthenticated(db *gorm.DB, r *http.Request) (bool, error) {
	_, err := Authenticate(db, r)
	if err != nil && !errors.IsRecordNotFound(err) {
		return false, err
	}

	return err != nil, nil
}
