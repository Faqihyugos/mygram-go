package handler

import (
	"net/http"

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
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	newSosmed, err := h.sosmedService.SaveSosmed(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Sosmed failed")
	}

	formatter := sosmed.FormatSosmedSave(newSosmed)
	c.JSON(http.StatusCreated, formatter)
}
