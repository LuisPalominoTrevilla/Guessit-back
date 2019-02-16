package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/LuisPalominoTrevilla/Guessit-back/routers"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is alive and kickin!")
}

func main() {
	fmt.Println(os.Getenv("SECRET_JWT_KEY"))

	r := routers.GetRouter()
	http.ListenAndServe(":5000", r)
}
