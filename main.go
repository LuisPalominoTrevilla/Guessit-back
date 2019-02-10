package main

import (
	"fmt"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is alive")
}

func main() {
	fmt.Println("Hola mundo")
	http.HandleFunc("/api", api)
	http.ListenAndServe(":5000", nil)
}
