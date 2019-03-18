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

func (db *UserDB) RegisterUser(name string, lastname string, gender string, email string, username string, password string) error {
	filter := bson.D{{"email", email}}

	userToRegister := models.User{
		Name:     name,
		Username: username,
		Image:    "",
		Email:    email,
		Gender:   gender,
		LastName: lastname,
		Password: password,
		Age:      21,
	}
	fmt.Println("Trying to find user", name, lastname, "in db")

	var foundUser models.User
	err := db.Get(filter, &foundUser)

	if err == mongo.ErrNoDocuments {
		res, err := db.Insert(userToRegister)
		if err != nil {
			return err
		}
		fmt.Println("Added user", name, lastname, "to database", res.InsertedID)
	} else if err == nil {
		fmt.Println("Seed user already created with id ", foundUser.ID)
	} else {
		return err
	}
	return nil
}
