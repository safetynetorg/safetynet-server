package main

import (
	"safetynet/internal/database"
	"safetynet/internal/devices"
	"safetynet/internal/keys"
	"safetynet/internal/server"
)

func main() {
	keys.Load()
	database.Connect()
	go devices.RemoveUninstalledDevices()
	server.Run()
}
