package database

import (
	"github.com/google/uuid"
)

type safetynet_device struct {
	Id  uuid.UUID `bson:"deviceid,omitempty"`
	Lat float64   `bson:"lat,omitempty"`
	Lon float64   `bson:"lon,omitempty"`
}
