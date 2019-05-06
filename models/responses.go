package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// PersonalDataResponse holds personal data from a user
type PersonalDataResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Image    string `json:"image"`
	Email    string `json:"email"`
	LastName string `json:"lastName"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
}

// AuthenticationResponse holds authentication information
type AuthenticationResponse struct {
	Token  string             `json:"token"`
	UserID primitive.ObjectID `json:"userId"`
}

// ImagesResponse holds an array of images as response
type ImagesResponse struct {
	Images []*Image `json:"images"`
}

// GuessResponse holds the result from the guess made by the user
type GuessResponse struct {
	Correct bool   `json:"correct"`
	Message string `json:"message"`
}

// ImageGuess holds information about guesses from an image
type ImageGuess struct {
	Quantity int `json:"quantity"`
	Correct  int `json:"correct"`
}

// StatisticalImage holds all statistics and information of a particular image
type StatisticalImage struct {
	ID                  primitive.ObjectID `json:"id"`
	URL                 string             `json:"url"`
	Age                 int                `json:"age"`
	RegisteredGuesses   *ImageGuess        `json:"registeredGuesses"`
	UnregisteredGuesses *ImageGuess        `json:"unregisteredGuesses"`
	CreatedAt           time.Time          `json:"createdAt"`
}

// UserImagesResponse holds an array of user images with statistical data
type UserImagesResponse struct {
	Images []*StatisticalImage `json:"images"`
}
