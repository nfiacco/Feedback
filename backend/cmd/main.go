package main

import (
	"feedback/internal/application"
	"feedback/internal/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func runServer(app *application.App) {
	// No HTTPS needed since TLS is terminated by Google Cloud Run
	r := mux.NewRouter()

	r.HandleFunc("/hello", app.Hello).Methods("GET")
	r.HandleFunc("/login", app.Login).Methods("POST")
	r.HandleFunc("/send", app.SendFeedback).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {

	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	app := application.InitApp(db)

	runServer(app)
}
