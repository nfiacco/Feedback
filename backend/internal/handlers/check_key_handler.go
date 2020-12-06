package handlers

import (
	"feedback/internal/errors"
	"feedback/internal/users"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CheckKey(env Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	feedbackKey, ok := vars["key"]
	if !ok {
		return fmt.Errorf("missing key from CheckKey request URL: %s", r.URL.RequestURI())
	}

	_, err := users.LoadByFeedbackKey(env.Db, feedbackKey)
	if err != nil && !errors.IsRecordNotFound(err) {
		return err
	} else if errors.IsRecordNotFound(err) {
		return errors.NotFound
	}

	return nil
}
