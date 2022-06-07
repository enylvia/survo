package migrations

import "time"

type Transaction struct {
	ID     int `gorm:"primaryKey; not null"`
	UserId int `gorm:"column:user_id; not null"`
	Amount int `gorm:"not null"`
	Status string `gorm:"type:varchar(100); nullable"`
	Type string `gorm:"type:varchar(100); nullable"`
	CreatedAt time.Time
	UpdatedAt time.Time
}