package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal(err)
	}
}
