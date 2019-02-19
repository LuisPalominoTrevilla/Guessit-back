package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/LuisPalominoTrevilla/Guessit-back/models"

	auth "github.com/LuisPalominoTrevilla/Guessit-back/authentication"
	"github.com/mongodb/mongo-go-driver/bson"

	database "github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// UserController wraps the UserDB inside the controller
type UserController struct {
	userDB *database.UserDB
}

// Get serves as a simple get request for the model User
func (controller *UserController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde usuario")
}

// Login serves as the endpoint to loggin the user
func (controller *UserController) Login(w http.ResponseWriter, r *http.Request) {
	type cred struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var credentials cred
	decoder := json.NewDecoder(r.Body)
	// Read credentials from request body
	err := decoder.Decode(&credentials)
	if err != nil {
		log.Fatal(err)
	}
	// Create bson document to filter in DB
	filter := bson.D{{"username", credentials.Username}, {"password", credentials.Password}}
	var loggedUser models.User
	// Find user and password in database
	err = controller.userDB.Get(filter, &loggedUser)
	if err != nil {
		w.WriteHeader(401)
		// TODO: Change this to return a JSON object
		fmt.Fprintf(w, "User or password incorrect")
		return
	}

	token, err := auth.GenerateJWT(&loggedUser)
	if err != nil {
		w.WriteHeader(500)
		// TODO: Change this to return a JSON object
		fmt.Fprintf(w, "Error generating jwt")
		return
	}
	type CustomResponse struct {
		Token    string             `json:"token"`
		Username string             `json:"username"`
		UserID   primitive.ObjectID `json:"userId"`
	}
	// TODO: Place line below inside middleware
	response := CustomResponse{
		Token:    token,
		Username: loggedUser.Username,
		UserID:   loggedUser.ID,
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// InitializeController initializes the routes
func (controller *UserController) InitializeController(r *mux.Router) {
	r.HandleFunc("/", controller.Get).Methods(http.MethodGet)
	r.HandleFunc("/Login", controller.Login).Methods(http.MethodPost)
}

// SetUserController creates the userController and wraps the user collection into UserDB
func SetUserController(r *mux.Router, db *mongo.Database) {
	user := database.UserDB{Users: db.Collection("users")}
	userController := UserController{userDB: &user}
	userController.InitializeController(r)
}
