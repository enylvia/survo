package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUserForm(input RegisterInput) (User, error)
	LoginUserForm(input LoginInput) (User, error)
	UpdateUserForm(inputID DetailUserInput, input UpdateInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUserForm(input RegisterInput) (User, error) {
	var user User

	user.FullName = input.FullName
	user.Email = input.Email
	user.Username = input.Username
	user.Occupation = input.Occupation
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user.Password = string(hashPassword)

	data, err := s.repository.Create(user)
	if err != nil {
		return data, err
	}
	return data, nil

}

func (s *service) LoginUserForm(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	find, err := s.repository.FindByEmail(email)

	if err != nil {
		return find, err
	}
	if find.Email == "" {
		return find, errors.New("user with that email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(find.Password), []byte(password))
	if err != nil {
		return find, errors.New("password is wrong")
	}
	return find, nil

}

func (s *service) UpdateUserForm(inputID DetailUserInput, input UpdateInput) (User, error) {
	user, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("user not found")
	}
	user.FullName = input.FullName
	user.Email = input.Email
	user.Username = input.Username
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user.Password = string(hashPassword)
	user.Phone = input.Phone
	user.Birthday = input.Birthday
	user.Image = input.Image

	newData , err := s.repository.Update(user)
	if err != nil {
		return newData, err
	}
	return newData, nil

}
