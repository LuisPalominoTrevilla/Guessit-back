package db

import (
	"context"

	"github.com/LuisPalominoTrevilla/Guessit-back/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type UserDB struct {
	Users *mongo.Collection
}

func (db *UserDB) Get(filter bson.D, result *models.User) error {
	return db.Users.FindOne(context.TODO(), filter).Decode(&result)
}

func (db *UserDB) Insert(user models.User) (*mongo.InsertOneResult, error) {
	return db.Users.InsertOne(context.TODO(), user)
}
