package handler

import (
	"net/http"

	"github.com/Faqihyugos/mygram-go/auth"
	"github.com/Faqihyugos/mygram-go/helper"
	"github.com/Faqihyugos/mygram-go/photo"
	"github.com/Faqihyugos/mygram-go/user"
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
	currentUser := c.MustGet("currentUser").(user.User)
	userID := int(currentUser.ID)
	var input photo.SavePhotoInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Photo failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPhoto, err := h.photoService.SavePhoto(userID, input)
	if err != nil {
		response := helper.ApiResponse("Photo failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := photo.FormatPhotoCreate(newPhoto)
	response := helper.ApiResponse("Photo has been save", http.StatusCreated, "succes", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h *photoHandler) GetAllPhoto(c *gin.Context) {
	photos, err := h.photoService.FindAllPhoto()
	if err != nil {
		response := helper.ApiResponse("Photo failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := photo.FormatPhotos(photos)
	response := helper.ApiResponse("List of photos", http.StatusOK, "succes", formatter)
	c.JSON(http.StatusOK, response)
}
