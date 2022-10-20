package photo

import (
	"time"

	"github.com/Faqihyugos/mygram-go/user"
)

type Photo struct {
	ID        int       `gorm:"primaryKey;column:id"`
	Title     string    `gorm:"column:title"`
	Caption   string    `gorm:"column:caption"`
	PhotoUrl  string    `gorm:"column:photo_url"`
	UserID    int       `gorm:"foreignKey:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	User      user.User
}
