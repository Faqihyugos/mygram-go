package sosmed

import "errors"

type Service interface {
	SaveSosmed(ID int, input SosmedInput) (Sosmed, error)
	FindAllSosmed() ([]Sosmed, error)
	UpdateSosmed(ID int, input SosmedInput) (Sosmed, error)
	DeleteSosmed(ID int) (Sosmed, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveSosmed(ID int, input SosmedInput) (Sosmed, error) {
	sosmed := Sosmed{}

	sosmed.Name = input.Name
	sosmed.SocialMediaUrl = input.SociallMediaUrl
	sosmed.UserID = ID

	NewSosmed, err := s.repository.Create(sosmed)
	if err != nil {
		return NewSosmed, err
	}
	return NewSosmed, nil
}

func (s *service) FindAllSosmed() ([]Sosmed, error) {
	socialmedias, err := s.repository.FindAll()
	if err != nil {
		return socialmedias, err
	}

	return socialmedias, nil
}

func (s *service) UpdateSosmed(ID int, input SosmedInput) (Sosmed, error) {
	sosmed, err := s.repository.FindByID(ID)
	if err != nil {
		return sosmed, err
	}

	if sosmed.UserID != input.User.ID {
		return sosmed, errors.New("Not an owner of the user")
	}

	sosmed.Name = input.Name
	sosmed.SocialMediaUrl = input.SociallMediaUrl

	updateSosmed, err := s.repository.Update(sosmed)
	if err != nil {
		return updateSosmed, err
	}

	return updateSosmed, nil
}

func (s *service) DeleteSosmed(ID int) (Sosmed, error) {
	sosmed, err := s.repository.FindByID(ID)
	if err != nil {
		return sosmed, err
	}

	deleteSosmed, err := s.repository.Delete(sosmed)
	if err != nil {
		return deleteSosmed, err
	}
	return deleteSosmed, nil
}
