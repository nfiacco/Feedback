package handlers

import (
	"encoding/json"
	"feedback/internal/users"
	"net/http"
)

type CheckSessionResponse struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	FeedbackKey string `json:"feedback_key"`
}

func CheckSession(env Env, w http.ResponseWriter, r *http.Request) error {
	if !env.Auth.IsAuthenticated {
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	userAndIdentity, err := users.LoadUserAndIdentityByID(env.Db, env.Auth.Session.UserID)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(CheckSessionResponse{
		FirstName:   userAndIdentity.FirstName,
		LastName:    userAndIdentity.LastName,
		FeedbackKey: userAndIdentity.FeedbackKey,
	})
	return nil
}
