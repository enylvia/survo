package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"survorest/survey"
)

type surveyHandler struct {
	surveyService survey.Service
}

func NewSurveyHandler(surveyService survey.Service) *surveyHandler {
	return &surveyHandler{surveyService: surveyService}
}

func (h *surveyHandler) IndexSurvey(c *gin.Context) {
	surveys, err := h.surveyService.GetSurveyList(0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	c.HTML(http.StatusOK, "survey_index.html", gin.H{"surveys": surveys})
}
