package handler

import (
	"net/http"
	"strconv"

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
	var input photo.PhotoInput

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

func (h *photoHandler) UpdatePhoto(c *gin.Context) {
	var input photo.UpdatePhotoInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Photo failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//get  id photo
	idString := c.Param("photoId")
	id, _ := strconv.Atoi(idString)

	//get current user
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	updatedPhoto, err := h.photoService.UpdatePhoto(id, input)

	if err != nil {
		response := helper.ApiResponse("Failed to update photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := photo.FormatPhotoUpdate(updatedPhoto)
	response := helper.ApiResponse("Success to update user", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *photoHandler) DeletePhoto(c *gin.Context) {
	idString := c.Param("photoId")
	id, _ := strconv.Atoi(idString)

	_, err := h.photoService.DeletePhoto(id)
	if err != nil {
		response := helper.ApiResponse("Failed to delete photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success to delete photo", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
