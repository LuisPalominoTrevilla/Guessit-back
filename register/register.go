package register

import (
	"fmt"

	"github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/LuisPalominoTrevilla/Guessit-back/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func registerUser(userDB db.UserDB, name string, lastname string, email string, username string, password string) error {
	filter := bson.D{{"email", email}}

	userToRegister := models.User{
		Name:     name,
		Username: username,
		Image:    "",
		Email:    email,
		Gender:   "male",
		LastName: "Palomino",
		Password: password,
		Age:      21,
	}
	fmt.Println("Trying to find user", name, lastname, "in db")

	var foundUser models.User
	err := userDB.Get(filter, &foundUser)

	if err == mongo.ErrNoDocuments {
		res, err := userDB.Insert(userToRegister)
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
