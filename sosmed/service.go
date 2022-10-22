package sosmed

type Service interface {
	SaveSosmed(ID int, input SosmedInput) (Sosmed, error)
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
