package main

import (
	"awsCloud/config"
	"awsCloud/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("AWS cloud storage ...")

	// load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// stabling the database
	_, dbErr := config.ConnectionDB()
	if dbErr != nil {
		panic("Error in stabling database connection")
	}

	// start server
	router := routes.RouterSetup()
	router.Run(":" + config.PORT)
}
