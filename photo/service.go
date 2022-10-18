package photo

type Service interface {
	SavePhoto(input SavePhotoInput) (Photo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SavePhoto(input SavePhotoInput) (Photo, error) {
	photo := Photo{}

	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = input.PhotoUrl

	// memanggil repo save
	NewPhoto, err := s.repository.Save(photo)
	if err != nil {
		return NewPhoto, err
	}
	return NewPhoto, nil
}
