package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"hospital-management/internal/models"
	"hospital-management/internal/services"
	"hospital-management/pkg/database"
)

func TestAuthService_Register(t *testing.T) {
	// Initialize test database
	database.InitDB()

	// Create auth service
	authService := services.NewAuthService()

	// Test user
	user := &models.User{
		Email:     "test@example.com",
		Password:  "password123",
		Role:      models.RoleReceptionist,
		FirstName: "Test",
		LastName:  "User",
	}

	// Test registration
	err := authService.Register(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)

	// Test duplicate email
	err = authService.Register(user)
	assert.Error(t, err)
}

func TestAuthService_Login(t *testing.T) {
	// Initialize test database
	database.InitDB()

	// Create auth service
	authService := services.NewAuthService()

	// Test user
	user := &models.User{
		Email:     "login@example.com",
		Password:  "password123",
		Role:      models.RoleDoctor,
		FirstName: "Login",
		LastName:  "Test",
	}

	// Register user
	err := authService.Register(user)
	assert.NoError(t, err)

	// Test successful login
	token, err := authService.Login(user.Email, "password123")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Test invalid credentials
	token, err = authService.Login(user.Email, "wrongpassword")
	assert.Error(t, err)
	assert.Empty(t, token)

	// Test non-existent user
	token, err = authService.Login("nonexistent@example.com", "password123")
	assert.Error(t, err)
	assert.Empty(t, token)
} 