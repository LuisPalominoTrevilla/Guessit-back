package models

import "github.com/mongodb/mongo-go-driver/bson/primitive"

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
