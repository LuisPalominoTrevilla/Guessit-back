package db

import (
	"context"
	"fmt"

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

// GetAll implements the find action in a model
func (db *UserDB) GetAll() ([]*models.User, error) {
	var results []*models.User

	cur, err := db.Users.Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Println("Error 1")
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("Error 2")
			return nil, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		fmt.Println("Error 3")
		return nil, err
	}

	cur.Close(context.TODO())

	return results, nil
}

// Insert implements the InsertOne action in a model
func (db *UserDB) Insert(user models.User) (*mongo.InsertOneResult, error) {
	return db.Users.InsertOne(context.TODO(), user)
}
