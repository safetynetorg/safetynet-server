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

type Contact struct {
	Id  string  `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
	Question string `bson:"question,omitempty"`
}