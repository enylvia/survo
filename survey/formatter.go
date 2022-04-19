package survey

import "strings"

type SurveyFormatter struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	Category string `json:"category"`
	Target int `json:"target"`
	Point int `json:"point"`
	Count int `json:"count"`
	Questions QuestionFormatter `json:"question"`
}

type QuestionFormatter struct {
	SurveyId int `json:"survey_id"`
	UserId int `json:"user_id"`
	SurveyQuestion string `json:"survey_question"`
	QuestionType string `json:"question_type"`
	OptionName []string `json:"option_name"`
}

func FormatSurvey(survey Survey) SurveyFormatter {
	formatter := SurveyFormatter{
		ID: int(survey.Id),
		Title: survey.Title,
		Summary: survey.Summary,
		Category: survey.Category,
		Target: survey.Target,
		Point: survey.Point,
		Count: survey.Count,
	}
	return formatter
}

func FormatSurveyList(surveys []Survey) []SurveyFormatter {
	surveysFormatter := []SurveyFormatter{}

	for _, survey := range surveys {
		surveyformatter := FormatSurvey(survey)
		surveysFormatter = append(surveysFormatter, surveyformatter)
	}
	return surveysFormatter
}

func FormatSurveyDetail(survey Survey) SurveyFormatter {
	surveyDetailFormatter := SurveyFormatter{}

	surveyDetailFormatter.ID = int(survey.Id)
	surveyDetailFormatter.Title = survey.Title
	surveyDetailFormatter.Summary = survey.Summary
	surveyDetailFormatter.Category = survey.Category
	surveyDetailFormatter.Target = survey.Target
	surveyDetailFormatter.Point = survey.Point
	surveyDetailFormatter.Count = survey.Count

	questions := survey.Questions
	questionFormatter := QuestionFormatter{}
	questionFormatter.SurveyId = int(survey.Id)
	questionFormatter.UserId = int(survey.UserId)
	questionFormatter.SurveyQuestion = questions.SurveyQuestion
	questionFormatter.QuestionType = questions.QuestionType
		var question_option []string
		for _, option := range strings.Split(questions.OptionName, ",") {
			question_option = append(question_option, strings.TrimSpace(option))
		}
		questionFormatter.OptionName = question_option

		surveyDetailFormatter.Questions = questionFormatter

	return surveyDetailFormatter
}