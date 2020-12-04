package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RunServer(db *gorm.DB) {
	// No HTTPS needed since TLS is terminated by Google Cloud Run
	r := mux.NewRouter()

	generateRoutes(db, r)

	r.Use(CORSMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func generateRoutes(db *gorm.DB, r *mux.Router) {
	for _, route := range Routes {
		withDb := WrapWithEnv(db, route.HandlerFunc)
		withErrorHandling := WrapWithErrorHandling(withDb)
		r.Handle(route.Pattern, withErrorHandling).Methods(route.Method, "OPTIONS")
	}
}
