package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Rate represents the model for rate
type Rate struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ImageID    primitive.ObjectID `json:"imageId"`
	FromAuth   bool               `json:"fromAuth"`
	GuessedAge int                `json:"guessedAge"`
}
