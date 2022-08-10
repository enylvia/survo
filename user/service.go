package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUserForm(input RegisterInput) (User, error)
	LoginUserForm(input LoginInput) (User, error)
	UpdateUserForm(inputID DetailUserInput, input UpdateInput) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByID(userID int) (User, error)
	UploadAvatar(userID int, filePath string) (User, error)
	GetAllUser() ([]User, error)
	DeleteUser(userID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) DeleteUser(userID int) error {
	if userID == 0 {
		return errors.New("user not found")
	}
	findUser, _ := s.repository.FindByID(userID)
	err := s.repository.Delete(int(findUser.Id))
	if err != nil {
		return err
	}
	return nil
}
func (s *service) GetAllUser() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}
func (s *service) RegisterUserForm(input RegisterInput) (User, error) {
	var user User

	user.FullName = input.FullName
	if input.Email == "" {
		return user, errors.New("email is required")
	}
	user.Email = input.Email
	user.Username = input.Username
	user.Occupation = input.Occupation
	user.IsAdmin = "user"
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashPassword)
	data, err := s.repository.Create(user)
	if err != nil {
		return data, err
	}
	var attrib Attribut
	attrib.UserId = uint(data.Id)
	attrib.PostedSurvey = 0
	attrib.TotalRespondent = 0
	attrib.ParticipateSurvey = 0
	attrib.IsPremium = false
	attrib.Balance = 0
	s.repository.CreateAttribut(attrib)

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
	if input.Password == "" {
		return user, errors.New("please input your password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, errors.New("password is incorrect")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user.Password = string(hashPassword)
	user.Phone = input.Phone
	user.Birthday = input.Birthday

	newData, err := s.repository.Update(user)
	if err != nil {
		return newData, err
	}
	return newData, nil

}
func (s *service) UploadAvatar(userID int, filePath string) (User, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("user not found")
	}
	user.Image = filePath
	newData, err := s.repository.Update(user)
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (s *service) GetUserByEmail(email string) (User, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.Email == "" {
		return user, errors.New("user with that email not found")
	}
	return user, nil
}
func (s *service) GetUserByID(userID int) (User, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}
