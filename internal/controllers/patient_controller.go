package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hospital-management/internal/models"
	"hospital-management/internal/services"
)

type PatientController struct {
	patientService *services.PatientService
}

func NewPatientController(patientService *services.PatientService) *PatientController {
	return &PatientController{
		patientService: patientService,
	}
}

func (c *PatientController) CreatePatient(ctx *gin.Context) {
	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the creator ID from the authenticated user
	userID, _ := ctx.Get("user_id")
	patient.CreatedBy = userID.(uint)

	if err := c.patientService.CreatePatient(&patient); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, patient)
}

func (c *PatientController) GetPatient(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	patient, err := c.patientService.GetPatient(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, patient)
}

func (c *PatientController) UpdatePatient(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the updater ID from the authenticated user
	userID, _ := ctx.Get("user_id")
	patient.UpdatedBy = userID.(uint)

	if err := c.patientService.UpdatePatient(uint(id), &patient); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Patient updated successfully"})
}

func (c *PatientController) DeletePatient(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	if err := c.patientService.DeletePatient(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

func (c *PatientController) ListPatients(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	patients, total, err := c.patientService.ListPatients(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  patients,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
} 