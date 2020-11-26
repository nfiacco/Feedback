package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "google.golang.org/api/idtoken"
)

const CLIENT_ID = "621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com"

type LoginRequest struct {
    IdToken string `json:"idtoken"`
}

func hello(w http.ResponseWriter, r *http.Request) {
    url := fmt.Sprintf("%v %v%v %v", r.Method, r.Host, r.URL, r.Proto)
    log.Print(url)
    fmt.Fprintf(w, "Hi there, I love go!")
}

func login(w http.ResponseWriter, r *http.Request) {
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
    log.Printf("got payload %+v", payload);
    return
}

func sendFeedback(w http.ResponseWriter, r *http.Request) {
    return
}

func main() {
    // No HTTPS needed since TLS is terminated by Google Cloud Run
	r := mux.NewRouter()
    r.HandleFunc("/hello", hello).Methods("GET")
    r.HandleFunc("/login", login).Methods("POST")
    r.HandleFunc("/send", sendFeedback).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", r))
}
