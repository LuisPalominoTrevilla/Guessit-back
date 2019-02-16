package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

type Model interface{}

type Database interface {
	Get(bson.D, *Model)
	Insert(Model)
}

var db *mongo.Database

func InitDb() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://mongodb:27017")
	if err != nil {
		log.Fatal("There was an error connecting to the database")
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Couldnt find a server ", err)
	}
	fmt.Println("Connected to MongoDB!")
	db = client.Database("guessit")
	return db
}
