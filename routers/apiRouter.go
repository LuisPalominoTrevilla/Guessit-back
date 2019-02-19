package routers

import (
	"fmt"
	"net/http"

	"github.com/LuisPalominoTrevilla/Guessit-back/controllers"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/gorilla/mux"
)

func handleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is alive")
}

// SetAPIRouter sets the API router and its subRouters
func SetAPIRouter(r *mux.Router, db *mongo.Database) {
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.StrictSlash(true)
	apiRouter.HandleFunc("/", handleAPI).Methods("GET")
	userRouter := apiRouter.PathPrefix("/User").Subrouter()
	controllers.SetUserController(userRouter, db)
}
