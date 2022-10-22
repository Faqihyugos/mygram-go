package photo

import (
	"time"

	"github.com/Faqihyugos/mygram-go/user"
)

type PhotoFormatterCreate struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoFormatter struct {
	ID        int                     `json:"id"`
	Title     string                  `json:"title"`
	Caption   string                  `json:"caption"`
	PhotoUrl  string                  `json:"photo_url"`
	UserID    int                     `json:"user_id"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
	User      user.UserPhotoFormatter `json:"user"`
}

type PhotoUpdateFormatter struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoCommentFormatter struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

func FormatPhotoCreate(photo Photo) PhotoFormatterCreate {
	formatter := PhotoFormatterCreate{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		UserID:   photo.UserID,
		PhotoUrl: photo.PhotoUrl,
	}
	return formatter
}

func FormatPhoto(photo Photo) PhotoFormatter {
	formatter := PhotoFormatter{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		UserID:    photo.UserID,
		PhotoUrl:  photo.PhotoUrl,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
		User: user.UserPhotoFormatter{
			Email:    photo.User.Email,
			Username: photo.User.Username,
		},
	}
	return formatter
}

func FormatPhotos(photos []Photo) []PhotoFormatter {
	if len(photos) == 0 {
		return []PhotoFormatter{}
	}

	var photoFormatter []PhotoFormatter
	for _, photo := range photos {
		photoFormatter = append(photoFormatter, FormatPhoto(photo))
	}

	return photoFormatter
}

func FormatPhotoUpdate(photo Photo) PhotoUpdateFormatter {
	formatter := PhotoUpdateFormatter{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		UserID:    photo.UserID,
		PhotoUrl:  photo.PhotoUrl,
		UpdatedAt: photo.UpdatedAt,
	}
	return formatter
}
