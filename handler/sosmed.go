package handler

import (
	"net/http"
	"strconv"

	"github.com/Faqihyugos/mygram-go/auth"
	"github.com/Faqihyugos/mygram-go/helper"
	"github.com/Faqihyugos/mygram-go/sosmed"
	"github.com/Faqihyugos/mygram-go/user"
	"github.com/gin-gonic/gin"
)

type sosmedHandler struct {
	sosmedService sosmed.Service
	authService   auth.Service
}

func NewSosmedHandler(sosmedService sosmed.Service, authService auth.Service) *sosmedHandler {
	return &sosmedHandler{sosmedService, authService}
}

func (h *sosmedHandler) CreateSosmed(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := int(currentUser.ID)
	input := sosmed.SosmedInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Create account social media failed",
			"message": errors,
		})
		return
	}

	newSosmed, err := h.sosmedService.SaveSosmed(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	formatter := sosmed.FormatSosmedSave(newSosmed)
	c.JSON(http.StatusCreated, formatter)
}

func (h *sosmedHandler) GetAllSosmed(c *gin.Context) {
	socialmedias, err := h.sosmedService.FindAllSosmed()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	formatter := sosmed.FormatSocialMedias(socialmedias)
	c.JSON(http.StatusOK, formatter)
}

func (h *sosmedHandler) UpdateSosmed(c *gin.Context) {
	var input sosmed.SosmedInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Failed to update social media",
			"message": errors,
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	// get current user
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	updateSosmed, err := h.sosmedService.UpdateSosmed(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to update social media",
			"message": err.Error(),
		})
		return
	}
	response := sosmed.FormatSosmedUpdate(updateSosmed)
	c.JSON(http.StatusOK, response)
}

func (h *sosmedHandler) DeleteSosmed(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// get current user
	_, errMessage := h.sosmedService.DeleteSosmed(id)
	if errMessage != nil {
		c.JSON(http.StatusBadRequest, "Failed to delete social media")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
