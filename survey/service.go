package survey

type Service interface {
	CreateSurveyForm(input CreateSurveyInput)(Survey, error)
	GetSurvey(id int)(Survey, error)
}
type service struct {
	repository repository
}

func NewService(repository repository) *service{
	return &service{
		repository: repository,
	}
}

func (s *service) CreateSurveyForm(input CreateSurveyInput) (Survey, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetSurvey(id int) (Survey, error) {
	//TODO implement me
	panic("implement me")
}
