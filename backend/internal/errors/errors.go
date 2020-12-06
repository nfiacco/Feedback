package errors

import (
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type HttpError struct {
	Code              int
	ClientVisibleData string
}

func (e HttpError) Error() string {
	return e.ClientVisibleData
}

var NotFound = HttpError{
	Code:              http.StatusNotFound,
	ClientVisibleData: http.StatusText(http.StatusNotFound),
}

func Wrap(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}

func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func IsCookieNotFound(err error) bool {
	return errors.Is(err, http.ErrNoCookie)
}
