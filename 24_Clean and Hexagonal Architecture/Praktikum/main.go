package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}
	e := echo.New()
	routes.NewRoute(e, db)

	e.Start(":8080")
}
