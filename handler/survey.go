package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"survorest/helper"
	"survorest/survey"
	"survorest/user"
)

type surveyHandler struct {
	surveyService survey.Service
}

func NewSurveyHandler(surveyService survey.Service) *surveyHandler {
	return &surveyHandler{surveyService}

}

func (h *surveyHandler) CreateSurvey(c *gin.Context) {
	var surveyInput survey.CreateSurveyInput
	currentUser := c.MustGet("claims").(user.User)
	surveyInput.UserId = currentUser.Id

	err := c.ShouldBindJSON(&surveyInput)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Invalid Input", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newSurvey, err := h.surveyService.CreateSurveyForm(surveyInput)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Error Creating Survey", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := survey.FormatSurvey(newSurvey)
	response := helper.ApiResponse("Create Survey Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
func (h *surveyHandler) SurveyList(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	surveys, err := h.surveyService.GetSurveyList(userID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Survey not found", http.StatusNotFound, "error", errorMessage)
		c.JSON(http.StatusNotFound, response)
		return
	}
	formatter := survey.FormatSurveyList(surveys)
	response := helper.ApiResponse("Successfully get list survey", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *surveyHandler) GetSurveyDetail(c *gin.Context) {
	var surveyID survey.SurveyDetailID

	err := c.ShouldBindUri(&surveyID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Invalid Input", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	surveyDetail, err := h.surveyService.GetSurveyDetail(surveyID.ID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Survey not found", http.StatusNotFound, "error", errorMessage)
		c.JSON(http.StatusNotFound, response)
		return
	}
	formatter := survey.FormatSurveyDetail(surveyDetail)
	response := helper.ApiResponse("Successfully get survey detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *surveyHandler) AnswerQuestion(c *gin.Context) {
	var answerSurvey []survey.AnswerInput

	err := c.ShouldBindJSON(&answerSurvey)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Invalid Input", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	answer, err := h.surveyService.AnswerQuestion(answerSurvey)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Error Answer Question", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := survey.FormatAnswer(answer)
	response := helper.ApiResponse("Successfully answer question", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}