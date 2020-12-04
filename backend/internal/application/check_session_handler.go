package application

import (
	"feedback/internal/authentication"
	"feedback/internal/errors"
	"net/http"
)

func (app *App) CheckSession(w http.ResponseWriter, r *http.Request) error {
	isAuthenticated, err := authentication.IsAuthenticated(app.db, r)
	if err != nil {
		return err
	}

	if !isAuthenticated {
		return errors.HttpError{Code: http.StatusUnauthorized, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
