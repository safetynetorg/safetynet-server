package env

import (
	"github.com/joho/godotenv"
)

func Load() {
	godotenv.Load(".env")
}
