package photo

import "time"

type PhotoFormatter struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatPhoto(photo Photo) PhotoFormatter {
	formatter := PhotoFormatter{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		UserID:   photo.UserID,
		PhotoUrl: photo.PhotoUrl,
	}
	return formatter
}
