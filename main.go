package main

import (
	"net/http"

	"github.com/LuisPalominoTrevilla/Guessit-back/boot/seeder"
	"github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/LuisPalominoTrevilla/Guessit-back/routers"
)

func main() {

	database := db.InitDb()
	seeder.SeedModels(database)

	r := routers.GetRouter(database)
	http.ListenAndServe(":5000", r)
}
