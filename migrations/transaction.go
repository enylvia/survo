package migrations

import "time"

type Transaction struct {
	ID     int `gorm:"primaryKey; not null"`
	UserID int `gorm:"not null"`
	Amount int `gorm:"not null"`
	Status string `gorm:"type:varchar(100); nullable"`
	Code   string `gorm:"type:varchar(100); nullable"`
	PaymentURL string `gorm:"type:varchar(242); nullable"`
	CreatedAt time.Time
	UpdatedAt time.Time
}