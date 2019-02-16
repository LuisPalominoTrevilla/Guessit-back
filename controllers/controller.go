package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	Get(http.ResponseWriter, *http.Request)
	InitializeController(*mux.Router)
}
