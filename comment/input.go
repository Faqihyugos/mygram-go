package comment

type CommentInput struct {
	Message string `json:"message" binding:"required"`
	PhotoID int    `json:"photo_id"`
}
