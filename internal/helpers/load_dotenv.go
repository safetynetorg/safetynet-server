package helpers

import (
	"github.com/joho/godotenv"
)

// load environment
func LoadDotEnv() {
	godotenv.Load(".env")
}
