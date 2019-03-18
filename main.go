package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LuisPalominoTrevilla/Guessit-back/boot/seeder"
	"github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/LuisPalominoTrevilla/Guessit-back/redis"
	"github.com/LuisPalominoTrevilla/Guessit-back/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Loaded .env successfully")
	database := db.InitDb()
	seeder.SeedModels(database)

	// TODO: Remove the next four lines and pass the client to the routers
	redisClient := redis.InitRedis()
	redisClient.SetArbitraryPair("lastname", "palomino")
	res, _ := redisClient.GetStringValue("lastname")
	fmt.Println(res)

	r := routers.GetRouter(database)
	http.ListenAndServe(":5000", r)
}
