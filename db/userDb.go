package db

import (
	"context"

	"github.com/LuisPalominoTrevilla/Guessit-back/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// UserDB serves as the wrapper for the users colelction
type UserDB struct {
	Users *mongo.Collection
}

// Get implements the findOne action in a model
func (db *UserDB) Get(filter bson.D, result *models.User) error {
	return db.Users.FindOne(context.TODO(), filter).Decode(&result)
}

// Insert implements the InsertOne action in a model
func (db *UserDB) Insert(user models.User) (*mongo.InsertOneResult, error) {
	return db.Users.InsertOne(context.TODO(), user)
}