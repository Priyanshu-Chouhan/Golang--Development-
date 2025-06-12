package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hospital-management/internal/controllers"
	"hospital-management/internal/routes"
	"hospital-management/internal/services"
	"hospital-management/pkg/database"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	database.InitDB()

	// Create services
	authService := services.NewAuthService()
	patientService := services.NewPatientService()

	// Create controllers
	authController := controllers.NewAuthController(authService)
	patientController := controllers.NewPatientController(patientService)

	// Initialize router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, authController, patientController)

	// Get port from environment variable or use default
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 