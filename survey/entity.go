package survey

import "time"

type Survey struct {
	Id         int64    `gorm:"primaryKey; not_null"`
	UserId     int64    `gorm:"column:user_id; not null"`
	Title      string   `gorm:"type:varchar(100); not null"`
	Summary    string   `gorm:"type:varchar(255); not null"`
	Target     int      `gorm:"not null"`
	Point      int      `gorm:"not null"`
	Question   Question `gorm:"ForeignKey:SurveyId"`
	Answer     Answer   `gorm:"ForeignKey:SurveyId"`
	Softdelete time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type Question struct {
	Id           int64  `gorm:"primaryKey; not_null"`
	SurveyId     int64  `gorm:"column:survey_id; not null"`
	UserId       int64  `gorm:"column:user_id; not null"`
	QuestionType string `gorm:"type:varchar(100); not null"`
	OptionName   string `gorm:"type:varchar(100);"`
	Answer       Answer `gorm:"ForeignKey:QuestionId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type Answer struct {
	Id         int64 `gorm:"primaryKey; not_null"`
	SurveyId   int64 `gorm:"column:survey_id; not null"`
	UserId     int64 `gorm:"column:user_id; not null"`
	QuestionId int64 `gorm:"column:question_id; not null"`
	Respond string 	 `gorm:"type:varchar(100)"`
}
