package main

import (
	"safetynet/internal/database"
	"safetynet/internal/devices"
	"safetynet/internal/keys"
	"safetynet/internal/server"
	"safetynet/internal/alert"
)

func main() {
	keys.Load()
	database.Connect()
	alert.InitClient()
	go devices.RemoveUninstalledDevices()
	server.Run()
}
