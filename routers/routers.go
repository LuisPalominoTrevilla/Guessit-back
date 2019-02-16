package routers

import (
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func GetRouter(db *mongo.Database) *mux.Router {
	r := mux.NewRouter()
	SetAPIRouter(r, db)
	return r
}
