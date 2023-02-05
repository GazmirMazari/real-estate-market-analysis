package routes

import (
	"net/http"

	"github.com/GazmirMazari/real-estate-market-analysis/routes/handler"
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
