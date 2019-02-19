package seeder

import (
	database "github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// SeedModels controlls the order in which models are seeded
func SeedModels(db *mongo.Database) {
	userDB := database.UserDB{Users: db.Collection("users")}
	seedUsers(userDB)
}
