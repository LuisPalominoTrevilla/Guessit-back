package db

import (
	"context"
	"fmt"

	"github.com/LuisPalominoTrevilla/Guessit-back/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// RateDB serves as a wrapper for the rate collection
type RateDB struct {
	Rates *mongo.Collection
}

// Get retrieves filtered rates from the database
func (db *RateDB) Get(filter bson.D) ([]*models.Rate, error) {
	cur, err := db.Rates.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var result []*models.Rate

	for cur.Next(context.TODO()) {
		var rate models.Rate
		err := cur.Decode(&rate)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		result = append(result, &rate)
	}
	return result, nil
}

// Insert implements the InsertOne action in a model
func (db *RateDB) Insert(rate models.Rate) (*mongo.InsertOneResult, error) {
	return db.Rates.InsertOne(context.TODO(), rate)
}
