package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ApiError
type ApiError struct {
	Message string `json:"message" example:"message"`
	Error   string `json:"error" example:"error"`
}

// ApiResponseError json response
func ApiResponseError(c *gin.Context, status int, message string, err interface{}) {
	var apiError ApiError
	apiError.Message = message
	apiError.Error = err.(string)
	c.JSON(status, apiError)
}

func FormatValidationError(err error) []string {
	var errors []string
	validationErrors := err.(validator.ValidationErrors)
	for _, e := range validationErrors {
		errors = append(errors, e.Error())
	}
	return errors
}
