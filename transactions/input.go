package transactions

type GetTransactionUserInput struct {
	ID int `uri:"id" binding:"required"`
}