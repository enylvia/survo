package user

import "time"

type User struct {
	Id         int64 `gorm:"primaryKey; not null"`
	FullName   string `gorm:"type:varchar(100); not null"`
	Email      string `gorm:"type:varchar(100); not null"`
	Username   string `gorm:"type:varchar(100); not null"`
	Occupation string `gorm:"type:varchar(100); not null"`
	Password   string `gorm:"type:varchar(100); not null"`
	Image      string `gorm:"type:varchar(100); nullable"`
	Phone      string `gorm:"type:varchar(100); nullable"`
	Birthday   string `gorm:"type:varchar(100); nullable"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

type Attribut struct {
	Id         	int64 `gorm:"primaryKey; not null"`
	UserId     	uint `gorm:"not null"`
	IsPremium 	bool `gorm:"not null"`
	Balance 	int `gorm:"not null"`
}