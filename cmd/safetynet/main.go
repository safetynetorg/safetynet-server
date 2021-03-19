package main

import (
	"safetynet/internal/database"
	"safetynet/internal/helpers"
	"safetynet/internal/server"
)

func main() {
	helpers.LoadDotEnv()
	database.Connect()

	server.Run()
}
