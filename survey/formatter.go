package survey

import "strings"

type SurveyFormatter struct {
	ID        int                 `json:"id"`
	UserId 	int `json:"user_id"`
	Title     string              `json:"title"`
	Summary   string              `json:"summary"`
	Category  string              `json:"category"`
	Target    int                 `json:"target"`
	Point     int                 `json:"point"`
	Count     int                 `json:"count"`
	Questions []QuestionFormatter `json:"question"`
}

type QuestionFormatter struct {
	ID             int      `json:"id"`
	SurveyId       int      `json:"survey_id"`
	UserId         int      `json:"user_id"`
	SurveyQuestion string   `json:"survey_question"`
	QuestionType   string   `json:"question_type"`
	OptionName     []string `json:"option_name"`
}
type AnswerFormatter struct {
	SurveyId   int    `json:"survey_id"`
	UserId     int    `json:"user_id"`
	QuestionId int    `json:"question_id"`
	Respond    string `json:"respond"`
}

func FormatSurvey(survey Survey) SurveyFormatter {
	formatter := SurveyFormatter{
		ID:       int(survey.Id),
		UserId:   int(survey.UserId),
		Title:    survey.Title,
		Summary:  survey.Summary,
		Category: survey.Category,
		Target:   survey.Target,
		Point:    survey.Point,
		Count:    survey.Count,
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
func FormatQuestion(questions Question) QuestionFormatter {
	questionsFormatter := QuestionFormatter{}
	questionsFormatter.ID = int(questions.Id)
	questionsFormatter.SurveyId = int(questions.SurveyId)
	questionsFormatter.UserId = int(questions.UserId)
	questionsFormatter.SurveyQuestion = questions.SurveyQuestion
	questionsFormatter.QuestionType = questions.QuestionType
	var optionName []string
	for _, option := range strings.Split(questions.OptionName, ",") {
		optionName = append(optionName, strings.TrimSpace(option))
	}
	questionsFormatter.OptionName = optionName

	return questionsFormatter
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
	questionFormatter := []QuestionFormatter{}
	for _, question := range questions {
		questionsFormatter := FormatQuestion(question)
		questionFormatter = append(questionFormatter, questionsFormatter)
	}
	surveyDetailFormatter.Questions = questionFormatter
	return surveyDetailFormatter
}
func FormatAnswer(answers Answer) AnswerFormatter {
	answerFormatter := AnswerFormatter{}
	answerFormatter.SurveyId = int(answers.SurveyId)
	answerFormatter.UserId = int(answers.UserId)
	answerFormatter.QuestionId = int(answers.QuestionId)
	answerFormatter.Respond = answers.Respond

	return answerFormatter
}
