package application

import (
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/api/idtoken"
)

const CLIENT_ID = "621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com"

type LoginRequest struct {
	IdToken string `json:"idtoken"`
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	var loginRequest LoginRequest
	err := decoder.Decode(&loginRequest)
	if err != nil {
		log.Printf("error parsing json %+v", err)
		return
	}

	validator, err := idtoken.NewValidator(ctx)
	if err != nil {
		log.Printf("error decoding token %+v", err)
		return
	}

	payload, err := validator.Validate(ctx, loginRequest.IdToken, CLIENT_ID)
	log.Printf("got payload subject: %+v", payload.Subject)
	log.Printf("got payload email: %+v", payload.Claims["email"])
	log.Printf("got payload family name: %+v", payload.Claims["family_name"])
	log.Printf("got payload given name: %+v", payload.Claims["given_name"])
	return
}
