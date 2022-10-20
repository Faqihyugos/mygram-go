package photo

import "github.com/Faqihyugos/mygram-go/user"

type PhotoInput struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

type UpdatePhotoInput struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	User     user.User
}
