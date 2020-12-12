package handlers

import (
	"encoding/json"
	"feedback/internal/config"
	"feedback/internal/emails"
	"feedback/internal/errors"
	"feedback/internal/sanitization"
	"feedback/internal/users"
	"html"
	"net/http"

	"github.com/mailgun/mailgun-go"
)

const mailgunDomain = "m.anonymousfeedback.app"

type SendRequest struct {
	FeedbackKey    string `json:"feedback_key"`
	EscapedContent string `json:"escaped_content"`
}

func SendFeedback(env Env, w http.ResponseWriter, r *http.Request) error {
	apiKey := config.GetMailgunApiKey()

	decoder := json.NewDecoder(r.Body)
	var sendRequest SendRequest
	err := decoder.Decode(&sendRequest)
	if err != nil {
		return err
	}

	user, err := users.LoadByFeedbackKey(env.Db, sendRequest.FeedbackKey)
	if err != nil && !errors.IsRecordNotFound(err) {
		return err
	}

	email, err := emails.LoadByUserID(env.Db, user.ID)
	if err != nil {
		return err
	}

	// Quill escapes the content already, so unescape it before sanitizing
	unescapedContent := html.UnescapeString(sendRequest.EscapedContent)
	textContent := sanitization.StrictPolicy.Sanitize(unescapedContent)
	sanitizedHtmlContent := sanitization.HtmlPolicy.Sanitize(unescapedContent)

	// TODO: add "don't reply to this email"?
	mg := mailgun.NewMailgun(mailgunDomain, apiKey)
	m := mg.NewMessage(
		"Anonymous Feedback <anonymous@anonymousfeedback.app>",
		"You have new feedback!",
		textContent,
		email.Email,
	)
	m.SetHtml(sanitizedHtmlContent)

	_, _, err = mg.Send(m)
	return err
}
