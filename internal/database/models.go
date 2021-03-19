package database

import (
	"github.com/google/uuid"
)

type safetynet_device_id struct {
	Id uuid.UUID `bson:"deviceid"`
}
