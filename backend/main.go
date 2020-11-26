package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love go!")
}

func login(w http.ResponseWriter, r *http.Request) {

    log.Info("hi");
    return
}

func sendFeedback(w http.ResponseWriter, r *http.Request) {
    return
}

func main() {
	r := mux.NewRouter()
    r.HandleFunc("/hello", hello).Methods("GET")
    r.HandleFunc("/login", login).Methods("POST").Schemes("https")
    r.HandleFunc("/send", sendFeedback).Methods("POST").Schemes("https")
    log.Fatal(http.ListenAndServe(":8080", r))
}
