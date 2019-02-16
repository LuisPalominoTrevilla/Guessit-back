package controllers

import (
	"fmt"
	"net/http"

	database "github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/LuisPalominoTrevilla/Guessit-back/models"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type UserController struct {
	userDB *database.UserDB
}

func (controller *UserController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde usuario")
}

func (controller *UserController) Login(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{"email", "luispalominot@hotmail.com"}}
	var foundUser models.User
	err := controller.userDB.Get(filter, &foundUser)
	if err != nil {
		fmt.Println("Not able to retrieve documents")
	}
	fmt.Println(foundUser.Name, foundUser.LastName)
	fmt.Fprintf(w, "The name is %s %s", foundUser.Name, foundUser.LastName)
}

func (controller *UserController) InitializeController(r *mux.Router) {
	r.HandleFunc("/", controller.Get).Methods(http.MethodGet)
	r.HandleFunc("/Login", controller.Login).Methods(http.MethodGet)
}

func SetUserController(r *mux.Router, db *mongo.Database) {
	user := database.UserDB{Users: db.Collection("users")}
	userController := UserController{userDB: &user}
	userController.InitializeController(r)
}
