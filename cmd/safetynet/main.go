package main

import (
	"os"
	"safetynet/internal/database"
	"safetynet/internal/helpers"
	"safetynet/internal/server"

	"github.com/kelvins/geocoder"
)

func main() {
	helpers.LoadDotEnv()
	geocoder.ApiKey = os.Getenv("GEOCODER_KEY")
	database.Connect()
	server.Run()
}
