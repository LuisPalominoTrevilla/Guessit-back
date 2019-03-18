package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

// Model serves as a wrapper for models
type Model interface{}

// Database is the interface that contains all relevant methods for each model
type Database interface {
	Get(bson.D, *Model)
	Insert(Model)
}

// InitDb initializes the database and returns it
func InitDb() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, os.Getenv("MONGO_HOST")+":"+os.Getenv("MONGO_PORT"))
	if err != nil {
		log.Fatal("There was an error connecting to the database")
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Couldnt find a server ", err)
	}
	fmt.Println("Connected to MongoDB!")
	return client.Database(os.Getenv("MONGO_DB"))
}
