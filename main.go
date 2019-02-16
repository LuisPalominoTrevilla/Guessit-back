package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is alive and kickin!")
}

func main() {
	fmt.Println(os.Getenv("SECRET_JWT_KEY"))
	r := mux.NewRouter()
	r.HandleFunc("/api", apiHandler)
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}
