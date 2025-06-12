package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FirstName    string    `gorm:"not null"`
	LastName     string    `gorm:"not null"`
	DateOfBirth  time.Time `gorm:"not null"`
	Gender       string    `gorm:"not null"`
	PhoneNumber  string    `gorm:"not null"`
	Email        string    `gorm:"uniqueIndex"`
	Address      string
	MedicalNotes string    `gorm:"type:text"`
	LastVisit    time.Time
	CreatedBy    uint      `gorm:"not null"` // ID of the receptionist who created the record
	UpdatedBy    uint      // ID of the doctor who last updated the record
} 