package routers

import (
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	SetAPIRouter(r)
	return r
}
