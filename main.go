package main

import (
	"github.com/kataras/iris/v12"
	"car-rentals-backend/repo"
	"car-rentals-backend/router"

)

func main() {
	// Connect Database
	db := repo.ConnectDatabase()
	defer db.Close()

	// Init Iris App
	app := iris.New()

	// Register router
	router.InitRouter(app)

	// Iris App listen port
	app.Listen(":8080")
}
