package models

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// User represents the model for the user
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Image    string             `json:"image"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	LastName string             `json:"lastName"`
	Password string             `json:"password"`
	Gender   string             `json:"gender"`
	Age      int                `json:"age"`
}
