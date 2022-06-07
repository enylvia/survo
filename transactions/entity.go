package transactions

import (
	"survorest/user"
	"time"
)

type Transaction struct {
	ID     int
	UserId int
	Amount int
	Status string
	Type string
	User user.User
	CreatedAt time.Time
	UpdatedAt time.Time
}
