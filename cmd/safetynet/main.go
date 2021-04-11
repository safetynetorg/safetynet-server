package main

import (
	"safetynet/internal/alert"
	"safetynet/internal/database"
	"safetynet/internal/devices"
	"safetynet/internal/helpers"
	"safetynet/internal/keys"
	"safetynet/internal/server"
)

func main() {
	keys.Load()
	helpers.AuthEmail()
	database.Connect()
	alert.InitClient()
	go devices.RemoveUninstalledDevices()
	server.Run()
}
