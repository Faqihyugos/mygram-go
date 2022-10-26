package handler

import (
	"net/http"
	"strconv"

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
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Create comment failed",
			"message": errors,
		})
		return
	}

	newComment, err := h.commentService.SaveComment(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	formatter := comment.FormatCommentCreate(newComment)
	c.JSON(http.StatusCreated, formatter)
}

func (h *commentHandler) GetAllComment(c *gin.Context) {
	comments, err := h.commentService.FindAllComment()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	formatter := comment.FormatComments(comments)
	c.JSON(http.StatusOK, formatter)
}

func (h *commentHandler) UpdateComment(c *gin.Context) {
	var input comment.UpdateCommentInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	// get id comment
	id, _ := strconv.Atoi(c.Param("id"))

	// get current user
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	updatedComment, err := h.commentService.UpdateComment(id, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to update comment",
			"message": err.Error(),
		})
		return
	}
	response := comment.FormatCommentUpdate(updatedComment)
	c.JSON(http.StatusOK, response)
}

func (h *commentHandler) DeleteComment(c *gin.Context) {
	// get id comment
	id, _ := strconv.Atoi(c.Param("id"))

	_, errMessage := h.commentService.DeleteComment(id)

	if errMessage != nil {
		c.JSON(http.StatusBadRequest, "Failed to delete comment")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Your comment has been successfully deleted"})
}
