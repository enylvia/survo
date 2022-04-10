package survey

import "gorm.io/gorm"

type Repository interface {
	GetSurvey(id string) (Survey, error)
	CreateSurvey(survey Survey) (Survey, error)
	CreateQuestion(question Question) (Question, error)
	CreateAnswer(answer []Answer) ([]Answer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}



func (r *repository) GetSurvey(id string) (Survey, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) CreateSurvey(survey Survey) (Survey, error) {
	err := r.db.Create(&survey).Error
	if err != nil {
		return survey, err
	}
	return survey,nil
}

func (r *repository) CreateQuestion(question Question) (Question, error) {
	err := r.db.Create(&question).Error
	if err != nil {
		return question, err
	}
	return question,err
}

func (r *repository) CreateAnswer(answer []Answer) ([]Answer, error) {
	//TODO implement me
	panic("implement me")
}
