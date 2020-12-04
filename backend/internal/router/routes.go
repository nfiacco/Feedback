package router

import (
	"net/http"

	"feedback/internal/handlers"
)

type EnvHandlerFunc func(http.ResponseWriter, *http.Request) error
type BaseHandlerFunc func(handlers.Env, http.ResponseWriter, *http.Request) error

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc BaseHandlerFunc
}

var Routes = []Route{
	Route{
		Name:        "Hello",
		Method:      "GET",
		Pattern:     "/hello",
		HandlerFunc: handlers.Hello,
	},
	Route{
		Name:        "Check session",
		Method:      "GET",
		Pattern:     "/check_session",
		HandlerFunc: handlers.CheckSession,
	},
	Route{
		Name:        "Login",
		Method:      "POST",
		Pattern:     "/login",
		HandlerFunc: handlers.Login,
	},
	Route{
		Name:        "Send feedback",
		Method:      "POST",
		Pattern:     "/send",
		HandlerFunc: handlers.SendFeedback,
	},
}
