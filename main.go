package main

import (
	"car-rentals-backend/repo"
	"car-rentals-backend/router"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func main() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

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
