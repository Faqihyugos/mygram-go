package handler

import (
	"net/http"

	"github.com/Faqihyugos/mygram-go/auth"
	"github.com/Faqihyugos/mygram-go/helper"
	"github.com/Faqihyugos/mygram-go/photo"
	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService photo.Service
	authService  auth.Service
}

func NewPhotoHandler(photoService photo.Service, authService auth.Service) *photoHandler {
	return &photoHandler{photoService, authService}
}

func (h *photoHandler) CreatePhoto(c *gin.Context) {
	var input photo.SavePhotoInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Photo failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPhoto, err := h.photoService.SavePhoto(input)
	if err != nil {
		response := helper.ApiResponse("Photo failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := photo.FormatPhoto(newPhoto)
	response := helper.ApiResponse("Photo has been save", http.StatusCreated, "succes", formatter)
	c.JSON(http.StatusOK, response)
}
