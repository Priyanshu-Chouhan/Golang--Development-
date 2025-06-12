package main

import (
	"log"

	"hospital-management/internal/models"
	"hospital-management/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	database.InitDB()

	// Auto migrate the schema
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Patient{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully")
}
