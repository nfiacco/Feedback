package application

import (
	"feedback/internal/errors"
	"log"
	"net/http"
	"strings"
)

type ApiHandlerFunc func(http.ResponseWriter, *http.Request) error

var ALLOWED_ORIGINS = []string{"https://anonymousfeedback.app", "https://www.anonymousfeedback.app"}
var ALLOWED_HEADERS = []string{"Content-Type"}

func WrapWithErrorHandling(handler ApiHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		log.Printf("Error: %+v", err)
		switch e := err.(type) {
		case errors.HttpError:
			http.Error(w, e.Error(), e.Code)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})
}

func isOriginAllowed(origin string) bool {
	if origin == "null" {
		return true
	}

	for _, allowedOrigin := range ALLOWED_ORIGINS {
		if origin == allowedOrigin {
			return true
		}
	}

	log.Printf("Origin not allowed: %s", origin)
	return false
}

// CORSMiddleware automatically sets the Access-Control-Allow-Origin and Access-Control-Allow-Credentials
// response headers for localhost:3000 (the frontend app)
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if isOriginAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", strings.Join(ALLOWED_HEADERS, ","))
		w.Header().Set("Access-Control-Allow-Methods", "POST")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
