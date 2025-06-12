package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleReceptionist Role = "receptionist"
	RoleDoctor      Role = "doctor"
)

type User struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      Role   `gorm:"type:varchar(20);not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	LastLogin time.Time
} 