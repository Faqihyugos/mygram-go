package handler

import (
	"net/http"

	"github.com/Faqihyugos/mygram-go/auth"
	"github.com/Faqihyugos/mygram-go/comment"
	"github.com/Faqihyugos/mygram-go/helper"
	"github.com/Faqihyugos/mygram-go/user"
	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService comment.Service
	authService    auth.Service
}

func NewCommentHandler(commentService comment.Service, authService auth.Service) *commentHandler {
	return &commentHandler{commentService, authService}
}

func (h *commentHandler) CreateComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := int(currentUser.ID)
	var input comment.CommentInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	newComment, err := h.commentService.SaveComment(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	formatter := comment.FormatCommentCreate(newComment)
	c.JSON(http.StatusCreated, formatter)
}
