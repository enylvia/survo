package transactions

type GetTransactionUserInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateTransactionInput struct {
	UserID int `json:"user_id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

type CreateTransactionPremium struct {
	ID int `uri:"id" binding:"required"`
}