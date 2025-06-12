package routes

import (
	"github.com/gin-gonic/gin"
	"hospital-management/internal/controllers"
	"hospital-management/internal/middleware"
	"hospital-management/internal/models"
)

func SetupRoutes(
	router *gin.Engine,
	authController *controllers.AuthController,
	patientController *controllers.PatientController,
) {
	// Public routes
	router.POST("/api/auth/login", authController.Login)
	router.POST("/api/auth/register", authController.Register)

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())

	// Patient routes
	patients := api.Group("/patients")
	{
		patients.GET("", patientController.ListPatients)
		patients.GET("/:id", patientController.GetPatient)

		// Receptionist only routes
		receptionist := patients.Group("")
		receptionist.Use(middleware.RoleMiddleware(models.RoleReceptionist))
		{
			receptionist.POST("", patientController.CreatePatient)
			receptionist.DELETE("/:id", patientController.DeletePatient)
		}

		// Doctor only routes
		doctor := patients.Group("")
		doctor.Use(middleware.RoleMiddleware(models.RoleDoctor))
		{
			doctor.PUT("/:id", patientController.UpdatePatient)
		}
	}
} 