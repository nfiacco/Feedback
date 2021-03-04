package handlers

import (
	"encoding/json"
	"feedback/internal/users"
	"net/http"
)

type CheckSessionResponse struct {
	FeedbackKey string `json:"feedback_key"`
}

func CheckSession(env Env, w http.ResponseWriter, r *http.Request) error {
	if !env.Auth.IsAuthenticated {
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	user, err := users.LoadUserByID(env.Db, env.Auth.Session.UserID)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(CheckSessionResponse{
		FeedbackKey: user.FeedbackKey,
	})
}
