package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetImagesRouter sets the API router and its subRouters
func SetImagesRouter(r *mux.Router) {
	fs := http.FileServer(http.Dir("static"))
	imagesRouter := r.PathPrefix("/images").Subrouter()
	imagesRouter.Handle("/{image}", http.StripPrefix("/images", fs))
	imagesRouter.Handle("/{directory}/{image}", http.StripPrefix("/images", fs))
}
