package user

import (
	"time"
)

type User struct {
	Id         uint
	FullName   string
	Email      string
	Username   string
	Occupation string
	Password   string
	Image      string
	Phone      string
	Birthday   string
	IsAdmin		string
	Attribut   Attribut
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Attribut struct {
	Id                uint
	UserId            uint
	PostedSurvey      int
	TotalRespondent   int
	ParticipateSurvey int
	IsPremium         bool
	Balance           int
}
