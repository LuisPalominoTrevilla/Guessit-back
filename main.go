package main

import (
	"net/http"

	"github.com/LuisPalominoTrevilla/Guessit-back/boot/seeder"
	"github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/LuisPalominoTrevilla/Guessit-back/redis"
	"github.com/LuisPalominoTrevilla/Guessit-back/routers"
)

func main() {

	database := db.InitDb()
	seeder.SeedModels(database)

	// TODO: Remove the next four lines and pass the client to the routers
	redisClient := redis.InitRedis()

	r := routers.GetRouter(database, redisClient)
	http.ListenAndServe(":5000", r)
}
