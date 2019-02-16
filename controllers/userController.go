package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// data should be changed later on to user model
type UserController struct {
	data string
}

func (controller *UserController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde usuario")
}

func (controller *UserController) initializeController(r *mux.Router) {
	r.HandleFunc("", controller.Get).Methods(http.MethodGet)
	r.HandleFunc("/", controller.Get).Methods(http.MethodGet)
}

func SetController(r *mux.Router) {
	userController := UserController{data: "nombre"}
	userController.initializeController(r)
}
