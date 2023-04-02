package main

import (
	"Praktikum/config"
	"Praktikum/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
