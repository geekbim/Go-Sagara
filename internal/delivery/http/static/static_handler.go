package static_handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StaticHandler(r *mux.Router) {
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
}
