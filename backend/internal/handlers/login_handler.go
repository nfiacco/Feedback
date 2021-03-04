package handlers

import (
	"context"
	"encoding/json"
	"feedback/internal/auth"
	"feedback/internal/errors"
	"feedback/internal/models"
	"feedback/internal/sessions"
	"feedback/internal/users"
	"feedback/internal/verifications"
	"net/http"

	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
)

const CLIENT_ID = "621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com"

type EmailAuthentication struct {
	Email          string `json:"email"`
	ValidationCode string `json:"validation_code"`
}

type LoginRequest struct {
	IdToken             *string              `json:"id_token,omitempty"`
	EmailAuthentication *EmailAuthentication `json:"email_authentication,omitempty"`
}

type LoginResponse struct {
	FeedbackKey string `json:"feedback_key"`
}

func Login(env Env, w http.ResponseWriter, r *http.Request) error {
	if env.Auth.IsAuthenticated {
		user, err := users.LoadUserByID(env.Db, env.Auth.Session.UserID)
		if err != nil {
			return err
		}

		return json.NewEncoder(w).Encode(LoginResponse{
			FeedbackKey: user.FeedbackKey,
		})
	}

	decoder := json.NewDecoder(r.Body)
	var loginRequest LoginRequest
	err := decoder.Decode(&loginRequest)
	if err != nil {
		return err
	}

	ctx := r.Context()

	var user *models.User
	var token *string
	switch {
	case loginRequest.IdToken != nil:
		user, token, err = googleLogin(ctx, env.Db, *loginRequest.IdToken)
	case loginRequest.EmailAuthentication != nil:
		user, token, err = emailLogin(env.Db, *loginRequest.EmailAuthentication)
	default:
		return errors.BadRequest
	}

	if err != nil {
		return err
	}

	auth.AddSessionCookie(w, *token)
	return json.NewEncoder(w).Encode(LoginResponse{
		FeedbackKey: user.FeedbackKey,
	})
}

func validateAndParseIdToken(ctx context.Context, idToken string) (*users.ExternalUserInfo, error) {
	validator, err := idtoken.NewValidator(ctx)
	if err != nil {
		return nil, err
	}

	payload, err := validator.Validate(ctx, idToken, CLIENT_ID)
	if err != nil {
		return nil, err
	}

	return &users.ExternalUserInfo{
		ExternalID: payload.Subject,
		Email:      payload.Claims["email"].(string),
		FirstName:  payload.Claims["family_name"].(string),
		LastName:   payload.Claims["given_name"].(string),
	}, nil
}

func googleLogin(ctx context.Context, db *gorm.DB, idToken string) (*models.User, *string, error) {
	externalUserInfo, err := validateAndParseIdToken(ctx, idToken)
	if err != nil {
		return nil, nil, err
	}

	user, err := users.GetOrCreateForExternalInfo(db, externalUserInfo)
	if err != nil {
		return nil, nil, err
	}

	token, err := sessions.Create(db, user.ID)
	if err != nil {
		return nil, nil, err
	}

	return user, token, nil
}

func emailLogin(db *gorm.DB, emailAuthentication EmailAuthentication) (*models.User, *string, error) {
	// the user is created when the validation code is sent
	user, err := users.LoadByEmail(db, emailAuthentication.Email)
	if err != nil {
		return nil, nil, err
	}

	_, err = verifications.VerifyCode(db, emailAuthentication.ValidationCode, user.ID)
	if err != nil {
		if errors.IsRecordNotFound(err) {
			return nil, nil, errors.Unauthorized
		}

		return nil, nil, err
	}

	token, err := sessions.Create(db, user.ID)
	if err != nil {
		return nil, nil, err
	}

	return user, token, nil
}
