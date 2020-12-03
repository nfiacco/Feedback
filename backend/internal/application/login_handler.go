package application

import (
	"context"
	"encoding/json"
	"feedback/internal/authentication"
	"feedback/internal/errors"
	"feedback/internal/sessions"
	"feedback/internal/user_identities"
	"feedback/internal/users"
	"log"
	"net/http"

	"google.golang.org/api/idtoken"
)

const CLIENT_ID = "621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com"
const SESSION_COOKIE_NAME = "X-Session-Token"

type LoginRequest struct {
	IdToken string `json:"idtoken"`
}

type LoginResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

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

func validateAndParseIdToken(ctx context.Context, loginRequest LoginRequest) (*users.ExternalUserInfo, error) {
	validator, err := idtoken.NewValidator(ctx)
	if err != nil {
		return nil, err
	}

	payload, err := validator.Validate(ctx, loginRequest.IdToken, CLIENT_ID)

	return &users.ExternalUserInfo{
		ExternalID: payload.Subject,
		Email:      payload.Claims["email"].(string),
		FirstName:  payload.Claims["family_name"].(string),
		LastName:   payload.Claims["given_name"].(string),
	}, nil
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) error {
	session, err := authentication.Authenticate(app.db, r)
	if err != nil && !errors.IsRecordNotFound(err) {
		return err
	}

	if err == nil {
		log.Printf("already authed")
		userIdentity, err := user_identities.LoadByUserID(app.db, session.UserID)
		if err != nil {
			return err
		}

		json.NewEncoder(w).Encode(LoginResponse{FirstName: userIdentity.FirstName, LastName: userIdentity.LastName})
		return nil
	}

	decoder := json.NewDecoder(r.Body)
	var loginRequest LoginRequest
	err = decoder.Decode(&loginRequest)
	if err != nil {
		return err
	}

	ctx := r.Context()
	externalUserInfo, err := validateAndParseIdToken(ctx, loginRequest)
	if err != nil {
		return err
	}

	user, err := users.GetOrCreateUser(app.db, externalUserInfo)
	if err != nil {
		return err
	}

	token, err := sessions.GenerateToken()
	if err != nil {
		return err
	}

	_, err = sessions.Create(app.db, *token, user.ID)
	if err != nil {
		return err
	}

	addCookie(w, SESSION_COOKIE_NAME, *token)
	json.NewEncoder(w).Encode(LoginResponse{FirstName: externalUserInfo.FirstName, LastName: externalUserInfo.LastName})
	return nil
}
