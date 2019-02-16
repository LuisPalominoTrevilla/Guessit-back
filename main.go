package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/LuisPalominoTrevilla/Guessit-back/boot/seeder"
	"github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/LuisPalominoTrevilla/Guessit-back/routers"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is alive and kickin!")
}

func main() {
	fmt.Println(os.Getenv("SECRET_JWT_KEY"))

	database := db.InitDb()
	seeder.SeedModels(database)

	r := routers.GetRouter(database)
	http.ListenAndServe(":5000", r)
}
