package sosmed

import (
	"time"

	"github.com/Faqihyugos/mygram-go/user"
)

type Sosmed struct {
	ID             int       `gorm:"primaryKey;column:id"`
	Name           string    `gorm:"not null;column:name"`
	SocialMediaUrl string    `gorm:"not null;column:social_media_url"`
	UserID         int       `grom:"foreignKey:user_id"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `grom:"column:updated_at"`
	User           user.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
