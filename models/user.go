package models

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	LastName string             `json:"lastName"`
	Password string             `json:"password"`
	Age      int                `json:"age"`
}
