package survey

type CreateSurveyInput struct {
	UserId            uint            `json:"user_id"`
	SurveyCategory    string          `json:"survey_category" binding:"required"`
	SurveyTitle       string          `json:"survey_title" binding:"required"`
	SurveyDescription string          `json:"survey_description" binding:"required"`
	Target            int             `json:"target" binding:"required"`
	Question          []QuestionInput `json:"question" binding:"required"`
}

type QuestionInput struct {
	SurveyId       uint   `json:"survey_id"`
	UserId         uint   `json:"user_id"`
	SurveyQuestion string `json:"survey_question" binding:"required"`
	QuestionType   string `json:"question_type" binding:"required"`
	OptionName     string `json:"option_name" binding:"required"`
}

type AnswerInput struct {
	Id uint `json:"id"`
	SurveyId uint `json:"survey_id"`
	UserId uint `json:"user_id"`
	QuestionId uint `json:"question_id"`
	Respond string `json:"respond"`
}

type SurveyDetailID struct {
	ID int `uri:"id" binding:"required"`
}
