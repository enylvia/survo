package handler

import (
	"net/http"
	"strconv"
	"survorest/survey"

	"github.com/gin-gonic/gin"
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
func (h *surveyHandler) DetailSurvey(c *gin.Context) {
	surveyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	surveys, err := h.surveyService.GetSurveyDetail(surveyId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	formatSurvey := survey.FormatSurveyDetail(surveys)
	c.HTML(http.StatusOK, "survey_detail.html", gin.H{"survey": formatSurvey})
}
func (h *surveyHandler) DeleteSurvey(c *gin.Context) {
	surveyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	err = h.surveyService.DeleteSurvey(surveyId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	c.Redirect(http.StatusMovedPermanently, "/surveys")
}
