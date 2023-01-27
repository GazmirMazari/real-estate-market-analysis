package handler

import (
	"net/http"
)

// Home serves as an example endpoint
// swagger:route POST / Home
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from a snippet box"))
}
