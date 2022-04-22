package survey

import (
	"errors"
)

type Service interface {
	CreateSurveyForm(survey CreateSurveyInput) (Survey, error)
	GetSurveyDetail(id int) (Survey, error)
	GetSurveyList(id int) ([]Survey, error)
	AnswerQuestion(input []AnswerInput) (Answer, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateSurveyForm(input CreateSurveyInput) (Survey, error) {
	var survey Survey
	survey.UserId = input.UserId
	survey.Title = input.SurveyTitle
	survey.Summary = input.SurveyDescription
	survey.Category = input.SurveyCategory
	survey.Target = input.Target
	survey.Point = 25

	survey, err := s.repository.CreateSurvey(survey)
	if err != nil {
		return survey, err
	}
	var questionInput Question
	for _, question := range input.Question {
		questionInput.SurveyId = survey.Id
		questionInput.UserId = input.UserId
		questionInput.SurveyQuestion = question.SurveyQuestion
		questionInput.QuestionType = question.QuestionType
		questionInput.OptionName = question.OptionName

		s.repository.CreateQuestion(questionInput)
	}
	return survey, nil
}

func (s *service) GetSurveyDetail(id int) (Survey, error) {
	survey, err := s.repository.GetSurveyDetail(id)
	if survey.Id != uint(id) {
		return survey, errors.New("Survey not found")
	}
	if err != nil {
		return survey, err
	}
	return survey, nil
}
func (s *service) GetSurveyList(id int) ([]Survey, error) {
	if id != 0 {
		survey, err := s.repository.GetSurveyByIDUser(id)
		if err != nil {
			return survey, err
		}
		return survey, nil
	}
	survey, err := s.repository.GetSurvey()
	if err != nil {
		return survey, err
	}
	return survey, nil
}

func (s *service) AnswerQuestion(input []AnswerInput) (Answer, error) {
	var answer Answer
	for _, val := range input {
		answer.UserId = val.UserId
		answer.SurveyId = val.SurveyId
		answer.QuestionId = val.QuestionId
		answer.Respond = val.Respond

		s.repository.CreateAnswer(answer)
	}
	return answer, nil
}
