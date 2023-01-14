package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ParkingSpot struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Zipcode    string             `json:"task,omitempty"`
	Avaliblity bool               `json:"status,omitempty"`
}
