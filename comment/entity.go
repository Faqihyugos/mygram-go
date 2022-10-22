package comment

import (
	"time"

	"github.com/Faqihyugos/mygram-go/photo"
	"github.com/Faqihyugos/mygram-go/user"
)

type Comment struct {
	ID        int         `gorm:"primaryKey;column:id"`
	UserID    int         `gorm:"foreignKey:user_id"`
	PhotoID   int         `gorm:"foreignKey:photo_id"`
	Message   string      `gorm:"not null;column:message"`
	CreatedAt time.Time   `gorm:"column:created_at"`
	UpdatedAt time.Time   `gorm:"column:updated_at"`
	User      user.User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photo     photo.Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
