package handlers

import (
	"context"
	"encoding/json"
	"feedback/internal/auth"
	"feedback/internal/sessions"
	"feedback/internal/users"
	"log"
	"net/http"

	"google.golang.org/api/idtoken"
)

const CLIENT_ID = "621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com"

type LoginRequest struct {
	IdToken string `json:"idtoken"`
}

type LoginResponse struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	FeedbackKey string `json:"feedback_key"`
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

func Login(env Env, w http.ResponseWriter, r *http.Request) error {
	if env.Auth.IsAuthenticated {
		log.Printf("already authed")
		userAndIdentity, err := users.LoadUserAndIdentityByID(env.Db, env.Auth.Session.UserID)
		if err != nil {
			return err
		}

		json.NewEncoder(w).Encode(LoginResponse{
			FirstName:   userAndIdentity.FirstName,
			LastName:    userAndIdentity.LastName,
			FeedbackKey: userAndIdentity.FeedbackKey,
		})
		return nil
	}

	decoder := json.NewDecoder(r.Body)
	var loginRequest LoginRequest
	err := decoder.Decode(&loginRequest)
	if err != nil {
		return err
	}

	ctx := r.Context()
	externalUserInfo, err := validateAndParseIdToken(ctx, loginRequest)
	if err != nil {
		return err
	}

	user, err := users.GetOrCreateUser(env.Db, externalUserInfo)
	if err != nil {
		return err
	}

	token, err := sessions.GenerateToken()
	if err != nil {
		return err
	}

	_, err = sessions.Create(env.Db, *token, user.ID)
	if err != nil {
		return err
	}

	auth.AddSessionCookie(w, *token)
	json.NewEncoder(w).Encode(LoginResponse{
		FirstName:   externalUserInfo.FirstName,
		LastName:    externalUserInfo.LastName,
		FeedbackKey: user.FeedbackKey,
	})
	return nil
}
