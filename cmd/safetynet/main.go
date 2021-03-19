package main

import (
	"fmt"
	"safetynet/internal/database"
	"safetynet/internal/helpers"
	"safetynet/internal/server"
)

func main() {
	helpers.LoadDotEnv()
	db := database.Connect()
	fmt.Println(db)

	server.Run()
}
