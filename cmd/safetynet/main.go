package main

import (
	"safetynet/internal/database"
	"safetynet/internal/keys"
	"safetynet/internal/server"
)

func main() {
	keys.Load()
	database.Connect()
	server.Run()
}
