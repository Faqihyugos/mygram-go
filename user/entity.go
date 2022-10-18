package user

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;column:id"`
	Username  string    `gorm:"unique;column:username"`
	Email     string    `gorm:"unique;column:email"`
	Password  string    `gorm:"column:password"`
	Age       uint      `gorm:"column:age"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
