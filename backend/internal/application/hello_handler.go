package application

import (
	"fmt"
	"log"
	"net/http"
)

func (app *App) Hello(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%v %v%v %v", r.Method, r.Host, r.URL, r.Proto)
	log.Print(url)
	fmt.Fprintf(w, "Hi there, I love go!")
}
