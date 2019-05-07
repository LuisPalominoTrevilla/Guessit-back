package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/LuisPalominoTrevilla/Guessit-back/models"
	"github.com/LuisPalominoTrevilla/Guessit-back/redis"

	auth "github.com/LuisPalominoTrevilla/Guessit-back/authentication"
	"github.com/mongodb/mongo-go-driver/bson"

	database "github.com/LuisPalominoTrevilla/Guessit-back/db"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// UserController wraps the UserDB inside the controller
type UserController struct {
	userDB         *database.UserDB
	redisClient    *redis.Client
	authMiddleware *auth.Middleware
}

// Get serves as a simple get request for the model User
func (controller *UserController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde usuario")
}

// Login godoc
// @Summary Login to system
// @Description login user to system
// @ID user-login
// @Accept  json
// @Produce  json
// @Param user body models.Credentials true "User credentials"
// @Success 200 {object} models.AuthenticationResponse
// @Failure 400 {string} Error message
// @Failure 401 {string} Error message
// @Failure 500 {string} Error message
// @Router /User/Login [post]
func (controller *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials
	decoder := json.NewDecoder(r.Body)
	// Read credentials from request body
	err := decoder.Decode(&credentials)
	if err != nil {
		fmt.Println("No hay cuerpo presente")
		w.WriteHeader(400)
		return
	}
	// Create bson document to filter in DB
	filter := bson.D{{"username", credentials.Username}, {"password", credentials.Password}}
	var loggedUser models.User
	// Find user and password in database
	err = controller.userDB.GetOne(filter, &loggedUser)
	if err != nil {
		w.WriteHeader(401)
		// TODO: Change this to return a JSON object
		fmt.Fprintf(w, "Usuario o contraseña incorrecta")
		return
	}

	token, err := auth.GenerateJWT(&loggedUser)
	if err != nil {
		w.WriteHeader(500)
		// TODO: Change this to return a JSON object
		fmt.Fprintf(w, "Error generando jwt")
		return
	}
	// TODO: Place line below inside middleware
	response := models.AuthenticationResponse{
		Token:  token,
		UserID: loggedUser.ID,
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// Logout godoc
// @Summary Logout user from system
// @Description Logout user
// @ID logout-endpoint
// @Produce plain
// @Security Bearer
// @Success 200 {string} OK
// @Failure 401 {string} Error message
// @Router /User/Logout [post]
func (controller *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	exp, err := strconv.ParseFloat(r.Header.Get("exp"), 64)
	var expInt int64
	if err != nil {
		expInt = time.Now().Add(time.Hour * 72).Unix()
	} else {
		expInt = int64(exp)
	}
	err = controller.redisClient.SetExpArbitraryPair("blacklist:"+r.Header.Get("token"), expInt, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Se cerró la sesión")
}

// PersonalData godoc
// @Summary Get personal data from a user
// @Description Retrieve personal data from user
// @ID personal-data-retrieval
// @Produce  json
// @Security Bearer
// @Success 200 {object} models.PersonalDataResponse
// @Failure 401 {string} Error message
// @Failure 500 {string} Error message
// @Router /User/PersonalData [get]
func (controller *UserController) PersonalData(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("uid")
	// Create bson document to filter in DB
	oid, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.D{{"_id", oid}}
	var user models.User
	err := controller.userDB.GetOne(filter, &user)
	if err != nil {
		w.WriteHeader(401)
		// TODO: Change this to return a JSON object
		fmt.Fprintf(w, "Error consiguiendo datos del usuario.")
		return
	}
	response := models.PersonalDataResponse{
		user.Name,
		user.Username,
		user.Image,
		user.Email,
		user.LastName,
		user.Gender,
		user.Age,
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// InitializeController initializes the routes
func (controller *UserController) InitializeController(r *mux.Router) {
	r.HandleFunc("/", controller.Get).Methods(http.MethodGet)
	r.HandleFunc("/Login", controller.Login).Methods(http.MethodPost)
	r.Handle("/PersonalData", controller.authMiddleware.AccessControl(controller.PersonalData)).Methods(http.MethodGet)
	r.HandleFunc("/Register", controller.Register).Methods(http.MethodPost)
	r.Handle("/Logout", controller.authMiddleware.AccessControl(controller.Logout)).Methods(http.MethodPost)
	r.HandleFunc("/Update", controller.UpdateProfile).Methods(http.MethodPatch)
}

// SetUserController creates the userController and wraps the user collection into UserDB
func SetUserController(r *mux.Router, db *mongo.Database, redisClient *redis.Client) {
	user := database.UserDB{Users: db.Collection("users")}
	userController := UserController{
		userDB:      &user,
		redisClient: redisClient,
		authMiddleware: &auth.Middleware{
			RedisClient: redisClient,
		},
	}
	userController.InitializeController(r)
}

// Register godoc
// @Summary Register
// @Description Register new user to the database
// @ID user-register
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 200 {string} success message
// @Failure 400 {string} error message
// @Failure 409 {string} error message
// @Failure 500 {string} error message
// @Router /User/Register [post]
func (controller *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)

	// Read credentials from request body
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("NO BODY PRESENT")
		w.WriteHeader(400)
		w.Write([]byte("JSON sin cuerpo"))
		return
	}
	// Create bson document to filter in DB
	filter1 := bson.D{{"email", user.Email}}

	var existing models.User
	cErr := controller.userDB.GetOne(filter1, &existing)

	if cErr == nil {
		w.WriteHeader(409)
		w.Write([]byte("Ya hay una cuenta vinculada a este correo."))
		return
	}

	filter2 := bson.D{{"username", user.Username}}

	cErr2 := controller.userDB.GetOne(filter2, &existing)

	if cErr2 == nil {
		w.WriteHeader(409)
		w.Write([]byte("Ese nombre de usuario ya está en uso."))
		return
	}

	userToRegister := models.User{
		Name:        user.Name,
		Username:    user.Username,
		Image:       user.Image,
		Email:       user.Email,
		Gender:      user.Gender,
		LastName:    user.LastName,
		Password:    user.Password,
		Age:         user.Age,
		RatedImages: []string{},
	}

	_, cErr3 := controller.userDB.Insert(userToRegister)

	if cErr3 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Estamos teniendo problemas, intenta de nuevo más tarde."))

		return
	}

}

// UpdateProfile godoc
// @Summary Update Profile
// @Description Allows to modify an user's registry in the database
// @ID update-profile
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 200 {string} success message
// @Failure 400 {string} error message
// @Failure 409 {string} error message
// @Failure 500 {string} error message
// @Router /User/Register [patch]
func (controller *UserController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)

	// Read credentials from request body
	err := decoder.Decode(&user)

	if err != nil {
		fmt.Println("NO BODY PRESENT")
		w.WriteHeader(400)
		w.Write([]byte("JSON sin cuerpo"))
		return
	}

	// Create bson document to find user that wants to update data in DB
	filter1 := bson.D{{"email", user.Email}}

	var existing models.User
	cErr := controller.userDB.GetOne(filter1, &existing)

	if cErr != nil {
		w.WriteHeader(409)
		w.Write([]byte("No existe una cuenta vinculada a este correo."))
		return
	}

	userToUpdate := models.User{
		Name:        existing.Name,
		Username:    existing.Username,
		Image:       existing.Image,
		Email:       existing.Email,
		Gender:      existing.Gender,
		LastName:    existing.LastName,
		Password:    existing.Password,
		Age:         existing.Age,
		RatedImages: existing.RatedImages,
	}

	if len(user.Username) > 0 {
		userToUpdate.Username = user.Username
	}

	if len(user.Image) > 0 {
		userToUpdate.Image = user.Image
	}

	if len(user.Name) > 0 {
		userToUpdate.Name = user.Name
	}

	if len(user.LastName) > 0 {
		userToUpdate.LastName = user.LastName
	}

	if len(user.Gender) > 0 {
		userToUpdate.Gender = user.Gender
	}

	if user.Age > 0 {
		userToUpdate.Age = user.Age
	}

	if len(user.Password) > 0 {
		userToUpdate.Password = user.Password
	}

	filterDoc := bson.D{{"email", existing.Email}}

	updateDoc := bson.D{{"$set", bson.D{{"name", userToUpdate.Name},
		{"image", userToUpdate.Image},
		{"username", userToUpdate.Username},
		{"email", userToUpdate.Email},
		{"lastname", userToUpdate.LastName},
		{"password", userToUpdate.Password},
		{"gender", userToUpdate.Gender},
		{"age", userToUpdate.Age}}}

	_, cErr2 := controller.userDB.UpdateOne(filterDoc, updateDoc)

	if cErr2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No se pudieron actualizar los datos. Intenta más tarde por favor."))
		fmt.Println(cErr2)
		return
	}

}
