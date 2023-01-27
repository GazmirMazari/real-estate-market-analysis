package routes

import (
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	h       *handler
}

type Routes []Route

func (r Route) ServeHTTP() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.handler.home)
}
