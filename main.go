package main

import (
	"log"
	"nami_navigation_log/database"
	"nami_navigation_log/models"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()

}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Entry{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
