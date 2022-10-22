package sosmed

import (
	"time"

	"github.com/Faqihyugos/mygram-go/user"
)

type SosmedFormatterSave struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	SociallMediaUrl string    `json:"social_media_url"`
	UserID          int       `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
}

type SosmedFormatter struct {
	ID              int                       `json:"id"`
	Name            string                    `json:"name"`
	SociallMediaUrl string                    `json:"social_media_url"`
	UserID          int                       `json:"user_id"`
	CreatedAt       time.Time                 `json:"created_at"`
	UpdatedAt       time.Time                 `json:"updated_at"`
	User            user.UserCommentFormatter `json:"user"`
}

type SosmedFormatterUpdate struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	SociallMediaUrl string    `json:"social_media_url"`
	UserID          int       `json:"user_id"`
	UpdatedAt       time.Time `json:"updated_at"`
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

func FormatSosmed(sosmed Sosmed) SosmedFormatter {
	formatter := SosmedFormatter{
		ID:              sosmed.ID,
		Name:            sosmed.Name,
		SociallMediaUrl: sosmed.SocialMediaUrl,
		UserID:          sosmed.UserID,
		CreatedAt:       sosmed.CreatedAt,
		UpdatedAt:       sosmed.UpdatedAt,
		User: user.UserCommentFormatter{
			ID:       sosmed.UserID,
			Email:    sosmed.User.Email,
			Username: sosmed.User.Username,
		},
	}

	return formatter
}

func FormatSocialMedias(socials []Sosmed) []SosmedFormatter {
	if len(socials) == 0 {
		return []SosmedFormatter{}
	}

	var socialmediasFormatter []SosmedFormatter
	for _, sosmed := range socials {
		socialmediasFormatter = append(socialmediasFormatter, FormatSosmed(sosmed))
	}

	return socialmediasFormatter
}

func FormatSosmedUpdate(sosmed Sosmed) SosmedFormatterUpdate {
	formatter := SosmedFormatterUpdate{
		ID:              sosmed.ID,
		Name:            sosmed.Name,
		SociallMediaUrl: sosmed.SocialMediaUrl,
		UserID:          sosmed.ID,
		UpdatedAt:       sosmed.UpdatedAt,
	}
	return formatter
}
