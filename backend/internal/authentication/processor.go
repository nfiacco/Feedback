package authentication

import (
	"feedback/internal/errors"
	"feedback/internal/models"
	"feedback/internal/sessions"
	"net/http"

	"gorm.io/gorm"
)

const SESSION_COOKIE_NAME = "X-Session-Token"

func addCookie(w http.ResponseWriter, name string, value string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, &cookie)
}

func AddSessionCookie(w http.ResponseWriter, token string) {
	addCookie(w, SESSION_COOKIE_NAME, token)
}

// todo: implement as middleware, pass auth object to handler. it should have "isAuthed: bool"
func Authenticate(db *gorm.DB, r *http.Request) (*models.Session, error) {
	cookie, err := r.Cookie(SESSION_COOKIE_NAME)
	if err != nil {
		return nil, err
	}

	session, err := sessions.LoadValidByToken(db, cookie.Value)
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
	if err != nil && !errors.IsRecordNotFound(err) && !errors.IsCookieNotFound(err) {
		return false, err
	}

	return err == nil, nil
}
