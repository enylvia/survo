package transactions

import (
	"errors"
	"survorest/user"
)

type service struct {
	repository     Repository
	userRepository user.Repository
}

type Service interface {
	GetAllTransaction() ([]Transaction, error)
	GetDataTransactionByIDUser(input GetTransactionUserInput) ([]Transaction, error)
	CreateTransactionWithdraw(input CreateTransactionInput) (Transaction, error)
	CreateTransactionPremium(input CreateTransactionPremium) (Transaction, error)
	ConfirmationTransaction(id int) (Transaction, error)
	DeclineTransaction(id int) (Transaction, error)
}

func NewService(repository Repository, userRepository user.Repository) *service {
	return &service{repository, userRepository}
}

func (s *service) GetAllTransaction() ([]Transaction, error) {
	transactions, err := s.repository.GetTransaction()
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetDataTransactionByIDUser(input GetTransactionUserInput) ([]Transaction, error) {
	if input.ID == 0 {
		return nil, errors.New("ID User is invalid")
	}
	transaction, err := s.repository.GetDataTransactionbyIDUser(input.ID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) CreateTransactionWithdraw(input CreateTransactionInput) (Transaction, error) {
	if input.UserID == 0 {
		return Transaction{}, errors.New("User ID is invalid")
	}
	findUser, _ := s.userRepository.FindByID(input.UserID)
	if findUser.Attribut.Balance < input.Amount {
		return Transaction{}, errors.New("Balance is not enough")
	}
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
	if input.ID == 0 {
		return Transaction{}, errors.New("ID User is invalid")
	}
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
	if id == 0 {
		return Transaction{}, errors.New("ID is Invalid")
	}
	var transaction Transaction
	findTransaction, err := s.repository.GetTransactionByID(id)
	if err != nil {
		return transaction, err
	}
	updateUser, _ := s.userRepository.FindByID(findTransaction.UserId)
	updateUser.Attribut.Balance = updateUser.Attribut.Balance - findTransaction.Amount
	s.userRepository.UpdateAttribut(updateUser.Attribut)

	findTransaction.Status = "Complete"

	update, err := s.repository.UpdateTransaction(findTransaction)
	if err != nil {
		return update, err
	}
	return update, nil

}
func (s *service) DeclineTransaction(id int) (Transaction, error) {
	if id == 0 {
		return Transaction{}, errors.New("ID is Invalid")
	}
	var transaction Transaction
	findTransaction, err := s.repository.GetTransactionByID(id)
	if err != nil {
		return transaction, err
	}
	findTransaction.Status = "Decline"
	update, err := s.repository.UpdateTransaction(findTransaction)
	if err != nil {
		return update, err
	}
	return update, nil
}
