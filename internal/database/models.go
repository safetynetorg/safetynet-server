package database

type SafetynetDevice struct {
	Id  string  `bson:"_id,omitempty"`
	Lat float64 `bson:"lat,omitempty"`
	Lon float64 `bson:"lon,omitempty"`
}

type Email struct {
	Id  string  `bson:"_id,omitempty"`
	Email string `bson:"email,omitempty"`
}