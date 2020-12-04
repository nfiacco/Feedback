package router

import (
	"feedback/internal/application"
	"feedback/internal/auth"
	"feedback/internal/errors"
	"feedback/internal/handlers"
	"log"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

var ALLOWED_ORIGINS = []string{"https://anonymousfeedback.app", "https://www.anonymousfeedback.app"}
var ALLOWED_HEADERS = []string{"Content-Type"}

func WrapWithEnv(db *gorm.DB, handler BaseHandlerFunc) EnvHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		auth, err := auth.GetAuthentication(db, r)
		if err != nil {
			return err
		}

		env := handlers.Env{
			Db:   db,
			Auth: *auth,
		}
		return handler(env, w, r)
	}
}

func WrapWithErrorHandling(handler EnvHandlerFunc) http.Handler {
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
	if !application.IsProd() && origin == "http://localhost:3000" {
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

// CORSMiddleware automatically sets the Access-Control-Allow-* response headers for the frontend app
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
