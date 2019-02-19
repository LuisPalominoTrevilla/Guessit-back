package routers

import (
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// GetRouter receives a database and returns the application router
func GetRouter(db *mongo.Database) *mux.Router {
	r := mux.NewRouter()
	SetAPIRouter(r, db)
	return r
}
