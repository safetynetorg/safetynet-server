package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SafetynetDevice struct {
	Id  primitive.ObjectID `bson:"_id,omitempty"`
	Lat float64            `bson:"lat,omitempty"`
	Lon float64            `bson:"lon,omitempty"`
}

type AlertThisId struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
}
