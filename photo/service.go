package photo

import "errors"

type Service interface {
	SavePhoto(ID int, input PhotoInput) (Photo, error)
	FindAllPhoto() ([]Photo, error)
	UpdatePhoto(ID int, input UpdatePhotoInput) (Photo, error)
	DeletePhoto(ID int) (Photo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SavePhoto(ID int, input PhotoInput) (Photo, error) {
	photo := Photo{}

	photo.UserID = ID
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

func (s *service) FindAllPhoto() ([]Photo, error) {
	photos, err := s.repository.FindAll()
	if err != nil {
		return photos, err
	}
	return photos, nil
}

func (s *service) UpdatePhoto(ID int, input UpdatePhotoInput) (Photo, error) {
	photo, err := s.repository.FindByID(ID)
	if err != nil {
		return photo, err
	}

	if photo.UserID != input.User.ID {
		return photo, errors.New("Not an owner of the user")
	}
	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = input.PhotoUrl

	updatedPhoto, err := s.repository.Update(photo)
	if err != nil {
		return updatedPhoto, err
	}
	return updatedPhoto, nil
}

func (s *service) DeletePhoto(ID int) (Photo, error) {
	photo, err := s.repository.FindByID(ID)
	if err != nil {
		return photo, err
	}

	deletedPhoto, err := s.repository.Delete(photo)
	if err != nil {
		return deletedPhoto, err
	}
	return deletedPhoto, nil
}
