package user

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
	//TODO implement me
	panic("implement me")
}

func (s *service) LoginUserForm(input LoginInput) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdateUserForm(input UpdateInput) (User, error) {
	//TODO implement me
	panic("implement me")
}
