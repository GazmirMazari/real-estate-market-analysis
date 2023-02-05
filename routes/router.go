package routes

import (
	"net/http"

	"real_estate_market/routes/handler"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
}

func InitializetRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home)
	return mux
}
