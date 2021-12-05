package main

import (
	"safetynet/internal/alert"
	"safetynet/internal/database"
	"safetynet/internal/devices"
	"safetynet/internal/env"
	"safetynet/internal/helpers"
	"safetynet/internal/server"
)

func main() {
	env.Load()
	helpers.AuthEmail()
	database.Connect()
	alert.InitClient()
	go devices.RemoveUninstalledDevices()
	server.Run()
}
