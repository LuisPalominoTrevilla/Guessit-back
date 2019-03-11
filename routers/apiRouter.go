package routers

import (
	"fmt"
	"net/http"

	"github.com/LuisPalominoTrevilla/Guessit-back/controllers"
	// Folder holding documentation
	_ "github.com/LuisPalominoTrevilla/Guessit-back/docs"
	"github.com/LuisPalominoTrevilla/Guessit-back/redis"
	"github.com/mongodb/mongo-go-driver/mongo"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// @title GuessIt API
// @version 0.1
// @description This is the API documentation for GuessIt
// @termsOfService http://swagger.io/terms/

// @contact.name Luis Palomino
// @contact.email luispalominot@hotmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:5000
// @BasePath /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func handleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is alive")
}

// SetAPIRouter sets the API router and its subRouters
func SetAPIRouter(r *mux.Router, db *mongo.Database, redisClient *redis.Client) {
	// Set swagger UI
	r.HandleFunc("/swagger/{rest}", httpSwagger.WrapHandler)
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.StrictSlash(true)
	apiRouter.HandleFunc("/", handleAPI).Methods("GET")
	userRouter := apiRouter.PathPrefix("/User").Subrouter()
	controllers.SetUserController(userRouter, db, redisClient)
}
