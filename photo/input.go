package photo

type SavePhotoInput struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	UserID   int    `json:"user_id,omitempty" swaggerignore:"true"`
}
