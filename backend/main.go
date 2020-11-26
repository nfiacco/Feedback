package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
    url := fmt.Sprintf("%v %v%v %v", r.Method, r.Host, r.URL, r.Proto)
    log.Print(url)
    fmt.Fprintf(w, "Hi there, I love go!")
}

func login(w http.ResponseWriter, r *http.Request) {

    log.Print("hi");
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
