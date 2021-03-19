package main

import (
	"fmt"
	"safetynet/internal/database"
	"safetynet/internal/server"
	"safetynet/internal/pkg/helpers"
)

func main() {
	helpers.LoadDotEnv()
	db := database.Connect()
	fmt.Println(db)

	server.Run()
}
