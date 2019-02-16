package seeder

import (
	"fmt"
	"log"

	"github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/LuisPalominoTrevilla/Guessit-back/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func seedUser(userDB db.UserDB, user models.User) error {
	filter := bson.D{{"email", user.Email}}
	fmt.Println("Trying to find user", user.Name, user.LastName, "in db")

	var foundUser models.User
	err := userDB.Get(filter, &foundUser)

	if err == mongo.ErrNoDocuments {
		res, err := userDB.Insert(user)
		if err != nil {
			return err
		}
		fmt.Println("Added user", user.Name, user.LastName, "to database", res.InsertedID)
	} else if err == nil {
		fmt.Println("Seed user already created with id ", foundUser.ID)
	} else {
		return err
	}
	return nil
}

func seedUsers(userDB db.UserDB) {
	luis := models.User{
		Name:     "Luis",
		Username: "luispalominot",
		Email:    "luispalominot@hotmail.com",
		LastName: "Palomino",
		Password: "palomitas123",
		Age:      21,
	}
	pietra := models.User{
		Name:     "Jorge",
		Username: "jorgePs",
		Email:    "jorgeps@gmail.com",
		LastName: "Pietra Santa",
		Password: "linux123",
		Age:      20,
	}
	dafne := models.User{
		Name:     "Dafne",
		Username: "daphnesabrina1",
		Email:    "daphnesabrina@gmail.com",
		LastName: "Medina",
		Password: "diego",
		Age:      22,
	}
	err := seedUser(userDB, luis)
	if err != nil {
		log.Fatal("User could not be instantiated")
	}
	err = seedUser(userDB, pietra)
	if err != nil {
		log.Fatal("User could not be instantiated")
	}
	err = seedUser(userDB, dafne)
	if err != nil {
		log.Fatal("User could not be instantiated")
	}
}
