package db

import (
	"context"

	"github.com/LuisPalominoTrevilla/Guessit-back/models"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// ImageDB serves as the wrapper for the image colelction
type ImageDB struct {
	Images *mongo.Collection
}

// Insert implements the InsertOne action in a model
func (db *ImageDB) Insert(image models.Image) (*mongo.InsertOneResult, error) {
	return db.Images.InsertOne(context.TODO(), image)
}
