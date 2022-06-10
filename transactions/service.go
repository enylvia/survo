package transactions

type service struct {
	repository Repository
}

type Service interface {
	GetAllTransaction() ([]Transaction, error)
	GetDataTransactionByIDUser(input GetTransactionUserInput) ([]Transaction, error)
	CreateTransactionWithdraw(input CreateTransactionInput) (Transaction, error)
	CreateTransactionPremium(input CreateTransactionPremium) (Transaction, error)
	ConfirmationTransaction(id int) (Transaction, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllTransaction() ([]Transaction, error) {
	transactions, err := s.repository.GetTransaction()
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetDataTransactionByIDUser(input GetTransactionUserInput) ([]Transaction, error) {
	transaction, err := s.repository.GetDataTransactionbyIDUser(input.ID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) CreateTransactionWithdraw(input CreateTransactionInput) (Transaction, error) {
	var transactions Transaction
	transactions.UserId = input.UserID
	transactions.Amount = input.Amount
	transactions.Status = "Pending"
	transactions.Type = "Withdraw"
	transaction, err := s.repository.CreateTransaction(transactions)

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) CreateTransactionPremium(input CreateTransactionPremium) (Transaction, error) {
	var transactions Transaction
	transactions.UserId = input.ID
	transactions.Amount = 35000
	transactions.Status = "Pending"
	transactions.Type = "Premium"

	transactions, err := s.repository.CreateTransaction(transactions)

	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) ConfirmationTransaction(id int) (Transaction, error) {
	var transaction Transaction
	findTransaction, err := s.repository.GetTransactionByID(id)
	if err != nil {
		return transaction, err
	}
	findTransaction.Status = "Complete"

	update,err := s.repository.UpdateTransaction(findTransaction)
	if err != nil {
		return update,err
	}
	return update,nil

}
