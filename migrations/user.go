package migrations

import (
	"survorest/survey"
	"survorest/transactions"
	"time"
)

type User struct {
	Id          uint                     `gorm:"primaryKey; not null"`
	FullName    string                   `gorm:"type:varchar(100); not null"`
	Email       string                   `gorm:"type:varchar(100); not null;uniqueIndex"`
	Username    string                   `gorm:"type:varchar(100); not null"`
	Occupation  string                   `gorm:"type:varchar(100); not null"`
	Password    string                   `gorm:"type:varchar(100); not null"`
	Image       string                   `gorm:"type:varchar(100); nullable"`
	Phone       string                   `gorm:"type:varchar(100); nullable"`
	Birthday    string                   `gorm:"type:varchar(100); nullable"`
	IsAdmin     string                   `gorm:"type:varchar(20); nullable"`
	Attribut    Attribut                 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignkey:UserId"`
	Survey      survey.Survey            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignkey:UserId"`
	Question    survey.Question          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignkey:UserId"`
	Answer      survey.Answer            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignkey:UserId"`
	Transaction transactions.Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignkey:UserId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Attribut struct {
	Id                uint `gorm:"primaryKey; not null, AUTO_INCREMENT"`
	UserId            uint `gorm:"column:user_id; not null"`
	PostedSurvey      int  `gorm:"not null"`
	TotalRespondent   int  `gorm:"not null"`
	ParticipateSurvey int  `gorm:"not null"`
	IsPremium         bool `gorm:"not null"`
	Balance           int  `gorm:"not null"`
}
