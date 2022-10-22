package comment

import (
	"time"

	"github.com/Faqihyugos/mygram-go/photo"
	"github.com/Faqihyugos/mygram-go/user"
)

type CommentFormatterCreate struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentFormatter struct {
	ID        int                         `json:"id"`
	Message   string                      `json:"message"`
	PhotoID   int                         `json:"photo_id"`
	UserID    int                         `json:"user_id"`
	CreatedAt time.Time                   `json:"created_at"`
	UpdatedAt time.Time                   `json:"updated_at"`
	User      user.UserCommentFormatter   `json:"user"`
	Photo     photo.PhotoCommentFormatter `json:"photo"`
}

type CommentFormatterUpdate struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatCommentCreate(comment Comment) CommentFormatterCreate {
	formatter := CommentFormatterCreate{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}
	return formatter
}

func FormatComment(comment Comment) CommentFormatter {
	formatter := CommentFormatter{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		User: user.UserCommentFormatter{
			ID:       comment.User.ID,
			Email:    comment.User.Email,
			Username: comment.User.Username,
		},
		Photo: photo.PhotoCommentFormatter{
			ID:       comment.Photo.ID,
			Title:    comment.Photo.Title,
			Caption:  comment.Photo.Caption,
			PhotoUrl: comment.Photo.PhotoUrl,
			UserID:   comment.Photo.UserID,
		},
	}
	return formatter
}

func FormatComments(comments []Comment) []CommentFormatter {
	if len(comments) == 0 {
		return []CommentFormatter{}
	}

	var commentsFormatter []CommentFormatter
	for _, comment := range comments {
		commentsFormatter = append(commentsFormatter, FormatComment(comment))
	}

	return commentsFormatter
}

func FormatCommentUpdate(comment Comment) CommentFormatterUpdate {
	formatter := CommentFormatterUpdate{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}
	return formatter
}
