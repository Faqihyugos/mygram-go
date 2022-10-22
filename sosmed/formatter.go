package sosmed

import "time"

type SosmedFormatterSave struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	SociallMediaUrl string    `json:"social_media_url"`
	UserID          int       `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
}

func FormatSosmedSave(sosmed Sosmed) SosmedFormatterSave {
	formatter := SosmedFormatterSave{
		ID:              sosmed.ID,
		Name:            sosmed.Name,
		SociallMediaUrl: sosmed.SocialMediaUrl,
		UserID:          sosmed.ID,
		CreatedAt:       sosmed.CreatedAt,
	}
	return formatter
}
