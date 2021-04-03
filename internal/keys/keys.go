package keys

import (
	"os"

	"github.com/joho/godotenv"
)

var MONGO_URL string
var SERVER_KEY string
var GEO_KEY string
var PORT string
var EMAIL_PASS string

func Load() {
	godotenv.Load("../../.env")
	MONGO_URL = os.Getenv("MONGO_URL")
	SERVER_KEY = os.Getenv("SERVER_KEY")
	GEO_KEY = os.Getenv("GEO_KEY")
	PORT = os.Getenv("PORT")
	EMAIL_PASS = os.Getenv("EMAIL_PASS")
}
