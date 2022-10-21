package comment

import "github.com/Faqihyugos/mygram-go/user"

type CommentInput struct {
	Message string `json:"message" binding:"required"`
	PhotoID int    `json:"photo_id"`
}

type UpdateCommentInput struct {
	Message string `json:"message" binding:"required"`
	User    user.User
}
