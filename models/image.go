package models

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// Image represents the model for the image
type Image struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	URL   string             `json:"url"`
	Age   int                `json:"age"`
	Owner primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
}
