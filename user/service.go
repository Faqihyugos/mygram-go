package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Age = input.Age
	user.Email = input.Email
	user.Username = input.Username

	// merubah input password menjadi hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	// menginput password yg sudah menjadi string ke struct entity user
	user.Password = string(passwordHash)

	// memanggil repo save
	NewUser, err := s.repository.Save(user)
	if err != nil {
		return NewUser, err
	}
	return NewUser, nil
}
