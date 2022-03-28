package survey

import "gorm.io/gorm"

type Repository interface {
	GetSurvey(id string) (Survey, error)
	CreateSurvey(survey Survey) (Survey, error)
	CreateQuestion(question []Question) ([]Question, error)
	CreateAnswer(answer []Answer) ([]Answer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}



func (r *repository) GetSurvey(id string) (Survey, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) CreateSurvey(survey Survey) (Survey, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) CreateQuestion(question []Question) ([]Question, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) CreateAnswer(answer []Answer) ([]Answer, error) {
	//TODO implement me
	panic("implement me")
}
