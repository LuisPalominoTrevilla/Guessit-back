package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetStaticRouter sets the static folder router
func SetStaticRouter(r *mux.Router) {
	fs := http.FileServer(http.Dir("static"))
	staticRouter := r.PathPrefix("/images").Subrouter()
	staticRouter.Handle("/{image}", http.StripPrefix("/images", fs))
	staticRouter.Handle("/{directory}/{image}", http.StripPrefix("/images", fs))
}
