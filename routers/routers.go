package routers

import (
	"net/http"

	"github.com/LuisPalominoTrevilla/Guessit-back/middleware"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// GetRouter receives a database and returns the application router
func GetRouter(db *mongo.Database) *mux.Router {
	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Use(middleware.Cors)
	SetAPIRouter(r, db)
	return r
}
