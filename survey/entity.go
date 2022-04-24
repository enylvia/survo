package survey

import (
	"survorest/user"
	"time"
)

type Survey struct {
	Id         uint
	UserId     uint
	Title      string
	Summary    string
	Category	string
	Target     int
	Point      int
	Questions   []Question
	Answer     Answer
	Count 	 int
	User user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type Question struct {
	Id           uint
	SurveyId     uint
	UserId       uint
	SurveyQuestion string
	QuestionType string
	OptionName   string
	Answer       []Answer
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type Answer struct {
	Id         uint
	SurveyId   uint
	UserId     uint
	QuestionId uint
	Respond string
}
