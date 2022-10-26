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
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Save photo failed",
			"message": errors,
		})
		return
	}

	newPhoto, err := h.photoService.SavePhoto(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	response := photo.FormatPhotoCreate(newPhoto)
	c.JSON(http.StatusCreated, response)
}

func (h *photoHandler) GetAllPhoto(c *gin.Context) {
	photos, err := h.photoService.FindAllPhoto()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	response := photo.FormatPhotos(photos)
	c.JSON(http.StatusOK, response)
}

func (h *photoHandler) UpdatePhoto(c *gin.Context) {
	var input photo.UpdatePhotoInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Failed to update photo",
			"message": errors,
		})
		return
	}

	//get  id photo
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	//get current user
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	updatedPhoto, err := h.photoService.UpdatePhoto(id, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	response := photo.FormatPhotoUpdate(updatedPhoto)
	c.JSON(http.StatusOK, response)
}

func (h *photoHandler) DeletePhoto(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	_, errMessage := h.photoService.DeletePhoto(id)
	if errMessage != nil {
		c.JSON(http.StatusBadRequest, "Failed to delete photo")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
