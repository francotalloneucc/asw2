package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Address     string             `bson:"address"`
	City        string             `bson:"city"`
	Country     string             `bson:"country"`
	Amenities   []string           `bson:"amenities"`
	Photos      []string           `bson:"photos"`
}
