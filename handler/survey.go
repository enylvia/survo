package handler

import "survorest/survey"

type surveyHandler struct {
	surveyService survey.Service
}

func NewSurveyHandler(surveyService survey.Service) *surveyHandler {
	return &surveyHandler{surveyService}
}