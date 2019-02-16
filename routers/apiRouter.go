package routers

import (
	"fmt"
	"net/http"

	"github.com/LuisPalominoTrevilla/Guessit-back/controllers"

	"github.com/gorilla/mux"
)

func handleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is alive")
}

func SetAPIRouter(r *mux.Router) {
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("", handleAPI).Methods("GET")
	apiRouter.HandleFunc("/", handleAPI).Methods("GET")
	userRouter := apiRouter.PathPrefix("/User").Subrouter()
	controllers.SetController(userRouter)
}
