package db

import (
	"context"
	"fmt"

	"github.com/LuisPalominoTrevilla/Guessit-back/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// ImageDB serves as the wrapper for the image colelction
type ImageDB struct {
	Images *mongo.Collection
}

// Get retrieves filtered models from the database
func (db *ImageDB) Get(filter bson.D) ([]*models.Image, error) {
	cur, err := db.Images.Find(context.TODO(), filter)
	fmt.Println("All good here")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var result []*models.Image

	for cur.Next(context.TODO()) {
		var image models.Image
		err := cur.Decode(&image)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		result = append(result, &image)
	}
	return result, nil
}

// Insert implements the InsertOne action in a model
func (db *ImageDB) Insert(image models.Image) (*mongo.InsertOneResult, error) {
	return db.Images.InsertOne(context.TODO(), image)
}
