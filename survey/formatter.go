package survey

type SurveyFormatter struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	Category string `json:"category"`
	Target int `json:"target"`
	Point int `json:"point"`
	Count int `json:"count"`
}

type QuestionFormatter struct {
	SurveyId int `json:"survey_id"`
	UserId int `json:"user_id"`
	SurveyQuestion string `json:"survey_question"`
	QuestionType string `json:"question_type"`
	OptionName string `json:"option_name"`
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