package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Image represents the model for the image
type Image struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	URL       string             `json:"url"`
	Age       int                `json:"age"`
	CreatedAt time.Time          `json:"createdAt"`
	Owner     primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
}
