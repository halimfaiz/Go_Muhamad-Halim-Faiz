package main

import (
	"Praktikum/config"
	"Praktikum/middlewares"
	"Praktikum/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
