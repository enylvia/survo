package migrations

import "time"

type Survey struct {
	Id        uint       `gorm:"primaryKey; not null, AUTO_INCREMENT"`
	UserId    uint       `gorm:"column:user_id; not null"`
	Title     string     `gorm:"type:varchar(100); not null"`
	Summary   string     `gorm:"type:varchar(255); not null"`
	Category  string     `gorm:"type:varchar(100); not null"`
	Target    int        `gorm:"not null"`
	Point     int        `gorm:"not null"`
	Questions []Question `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignkey:survey_id"`
	Answer    Answer     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignkey:survey_id"`
	Count     int        `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Question struct {
	Id             uint     `gorm:"primaryKey; not null"`
	SurveyId       uint     `gorm:"column:survey_id; not null"`
	UserId         uint     `gorm:"column:user_id; not null"`
	SurveyQuestion string   `gorm:"type:varchar(255); not null"`
	QuestionType   string   `gorm:"type:varchar(100); not null"`
	OptionName     string   `gorm:"type:varchar(100);"`
	Answer         []Answer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignkey:question_id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type Answer struct {
	Id         uint   `gorm:"primaryKey; not null"`
	SurveyId   uint   `gorm:"column:survey_id; not null"`
	UserId     uint   `gorm:"column:user_id; not null"`
	QuestionId uint   `gorm:"column:question_id; not null"`
	Respond    string `gorm:"type:varchar(100)"`
}
