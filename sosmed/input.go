package sosmed

type SosmedInput struct {
	Name            string `json:"name" binding:"required"`
	SociallMediaUrl string `json:"social_media_url" binding:"required"`
}
