package services

import (
	"errors"

	"hospital-management/internal/models"
	"hospital-management/pkg/database"
)

type PatientService struct{}

func NewPatientService() *PatientService {
	return &PatientService{}
}

func (s *PatientService) CreatePatient(patient *models.Patient) error {
	return database.DB.Create(patient).Error
}

func (s *PatientService) GetPatient(id uint) (*models.Patient, error) {
	var patient models.Patient
	if err := database.DB.First(&patient, id).Error; err != nil {
		return nil, errors.New("patient not found")
	}
	return &patient, nil
}

func (s *PatientService) UpdatePatient(id uint, patient *models.Patient) error {
	var existingPatient models.Patient
	if err := database.DB.First(&existingPatient, id).Error; err != nil {
		return errors.New("patient not found")
	}

	// Update fields
	existingPatient.FirstName = patient.FirstName
	existingPatient.LastName = patient.LastName
	existingPatient.DateOfBirth = patient.DateOfBirth
	existingPatient.Gender = patient.Gender
	existingPatient.PhoneNumber = patient.PhoneNumber
	existingPatient.Email = patient.Email
	existingPatient.Address = patient.Address
	existingPatient.MedicalNotes = patient.MedicalNotes
	existingPatient.UpdatedBy = patient.UpdatedBy

	return database.DB.Save(&existingPatient).Error
}

func (s *PatientService) DeletePatient(id uint) error {
	return database.DB.Delete(&models.Patient{}, id).Error
}

func (s *PatientService) ListPatients(page, pageSize int) ([]models.Patient, int64, error) {
	var patients []models.Patient
	var total int64

	offset := (page - 1) * pageSize

	if err := database.DB.Model(&models.Patient{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := database.DB.Offset(offset).Limit(pageSize).Find(&patients).Error; err != nil {
		return nil, 0, err
	}

	return patients, total, nil
} 