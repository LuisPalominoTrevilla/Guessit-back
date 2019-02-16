package seeder

import (
	database "github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func SeedModels(db *mongo.Database) {
	userDB := database.UserDB{Users: db.Collection("users")}
	seedUsers(userDB)
}
