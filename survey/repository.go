package survey

import "gorm.io/gorm"

type Repository interface {
	GetSurvey() ([]Survey, error)
	GetSurveyByIDUser(id int) ([]Survey, error)
	CreateSurvey(survey Survey) (Survey, error)
	CreateQuestion(question Question) (Question, error)
	CreateAnswer(answer []Answer) ([]Answer, error)
	GetSurveyDetail(id int)(Survey, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}



func (r *repository) GetSurvey() ([]Survey, error) {
	//TODO implement me
	var surveys []Survey
	err := r.db.Find(&surveys).Error
	if err != nil {
		return surveys, err
	}
	return surveys, nil
}

func (r *repository) GetSurveyByIDUser(id int) ([]Survey, error) {
	var surveys []Survey
	err := r.db.Where("user_id = ?", id).Find(&surveys).Error
	if err != nil {
		return surveys, err
	}
	return surveys, nil
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

func (r *repository) GetSurveyDetail(id int)(Survey, error){
	var surveys Survey
	err := r.db.Preload("Questions").Where("id = ?", id).Find(&surveys).Error
	if err != nil {
		return surveys, err
	}
	return surveys, nil
}

func (r *repository) CreateAnswer(answer []Answer) ([]Answer, error) {
	//TODO implement me
	panic("implement me")
}
