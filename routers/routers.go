package routers

import (
	"net/http"

	"github.com/LuisPalominoTrevilla/Guessit-back/middleware"
	"github.com/LuisPalominoTrevilla/Guessit-back/redis"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// GetRouter receives a database and returns the application router
func GetRouter(db *mongo.Database, redisClient *redis.Client) *mux.Router {
	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Use(middleware.Cors)
	SetAPIRouter(r, db, redisClient)
	SetStaticRouter(r)
	return r
}
