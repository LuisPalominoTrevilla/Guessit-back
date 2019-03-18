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

// Insert implements the InsertOne action in a model
func (db *UserDB) Insert(user models.User) (*mongo.InsertOneResult, error) {
	return db.Users.InsertOne(context.TODO(), user)
}

// Create enables the capability to create a new user
func (db *UserDB) Create(user models.User, filter bson.D) error {

	fmt.Println("Verifying if email", user.Email, "is not in db")

	var foundUser models.User
	err := db.Get(filter, &foundUser)

	if err == mongo.ErrNoDocuments {
		res, err := db.Insert(user)
		if err != nil {
			return err
		}
		fmt.Println("Added user", user.Name, user.LastName, "to database", res.InsertedID)
	} else if err == nil {
		fmt.Println("User already created with email ", foundUser.Email)
	} else {
		return err
	}
	return nil
}
