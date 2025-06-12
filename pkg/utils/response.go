package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SendSuccess(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func SendError(c *gin.Context, status int, message string, err error) {
	c.JSON(status, Response{
		Status:  status,
		Message: message,
		Error:   err.Error(),
	})
}

func SendValidationError(c *gin.Context, message string) {
	SendError(c, http.StatusBadRequest, message, nil)
}

func SendNotFound(c *gin.Context, message string) {
	SendError(c, http.StatusNotFound, message, nil)
}

func SendUnauthorized(c *gin.Context, message string) {
	SendError(c, http.StatusUnauthorized, message, nil)
}

func SendForbidden(c *gin.Context, message string) {
	SendError(c, http.StatusForbidden, message, nil)
}

func SendInternalServerError(c *gin.Context, err error) {
	SendError(c, http.StatusInternalServerError, "Internal server error", err)
} 