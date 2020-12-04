package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func Hello(env Env, w http.ResponseWriter, r *http.Request) error {
	url := fmt.Sprintf("%v %v%v %v", r.Method, r.Host, r.URL, r.Proto)
	log.Print(url)
	fmt.Fprintf(w, "Hi there, I love go!")

	return nil
}
