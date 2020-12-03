package application

import (
	"feedback/internal/errors"
	"log"
	"net/http"
)

type ApiHandlerFunc func(http.ResponseWriter, *http.Request) error

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

// CORSMethodMiddleware automatically sets the Access-Control-Allow-Origin and Access-Control-Allow-Credentials
// response headers for localhost:3000 (the frontend app)
func DevCORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "http://localhost:3000")

		next.ServeHTTP(w, req)
	})
}
