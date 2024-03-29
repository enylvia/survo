package survey

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetSurvey() ([]Survey, error)
	GetSurveyByIDUser(id int) ([]Survey, error)
	CreateSurvey(survey Survey) (Survey, error)
	CreateQuestion(question Question) (Question, error)
	CreateAnswer(answer Answer) (Answer, error)
	GetSurveyDetail(id int) (Survey, error)
	GetSurveyRespond(id int) ([]Answer, error)
	DeleteSurvey(id int) error
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
	return survey, nil
}

func (r *repository) CreateQuestion(question Question) (Question, error) {
	err := r.db.Create(&question).Error
	if err != nil {
		return question, err
	}
	return question, err
}

func (r *repository) GetSurveyDetail(id int) (Survey, error) {
	var surveys Survey
	err := r.db.Preload("Questions").Where("id = ?", id).Find(&surveys).Error
	if err != nil {
		return surveys, err
	}
	return surveys, nil
}

func (r *repository) CreateAnswer(answer Answer) (Answer, error) {
	err := r.db.Create(&answer).Error
	if err != nil {
		return answer, err
	}
	return answer, err
}

func (r *repository) GetSurveyRespond(id int) ([]Answer, error) {
	var respond []Answer
	err := r.db.Where("survey_id = ?", id).Find(&respond).Error
	if err != nil {
		return respond, err
	}
	return respond, nil
}

func (r *repository) DeleteSurvey(id int) error {
	err := r.db.Where("id = ?", id).Delete(&Survey{}).Error
	if err != nil {
		return err
	}
	return nil
}
