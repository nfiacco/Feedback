package handlers

import (
	"encoding/json"
	"feedback/internal/config"
	"feedback/internal/errors"
	"feedback/internal/users"
	"feedback/internal/verifications"
	"net/http"
	"regexp"

	"github.com/mailgun/mailgun-go"
)

const emailRegex = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

type ValidationCodeRequest struct {
	Email string `json:"email"`
}

func ValidationCode(env Env, w http.ResponseWriter, r *http.Request) error {
	apiKey := config.GetMailgunApiKey()

	decoder := json.NewDecoder(r.Body)
	var request ValidationCodeRequest
	err := decoder.Decode(&request)
	if err != nil {
		return err
	}

	// regex isn't perfect, but try not to send to poorly formatted emails
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(request.Email) {
		return errors.BadRequest
	}

	user, err := users.GetOrCreateForEmail(env.Db, request.Email)
	if err != nil {
		return err
	}

	code, err := verifications.Create(env.Db, user.ID)
	if err != nil {
		return err
	}

	mg := mailgun.NewMailgun(mailgunDomain, apiKey)
	m := mg.NewMessage(
		"Anonymous Feedback <anonymous@anonymousfeedback.app>",
		"Verification Code",
		*code,
		request.Email,
	)

	_, _, err = mg.Send(m)
	return err
}
