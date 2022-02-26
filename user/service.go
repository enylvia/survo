package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUserForm(input RegisterInput) (User, error)
	LoginUserForm(input LoginInput) (User, error)
	UpdateUserForm(input UpdateInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository)*service{
	return &service{repository}
}

func (s *service) RegisterUserForm(input RegisterInput) (User, error) {
	var user User

	hashPassword,_ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	user.FullName = input.FullName
	user.Email = input.Email
	user.Username = input.Username
	user.Occupation = input.Occupation
	user.Password = string(hashPassword)

	newUser,err := s.repository.Create(user)

	if err != nil {
		return newUser, err
	}
	return newUser, nil

}

func (s *service) LoginUserForm(input LoginInput) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateUserForm(input UpdateInput) (User, error) {
	//TODO implement me
	panic("implement me")
}
