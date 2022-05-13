package transactions

type service struct {
	repository Repository
}

type Service interface {
	GetAllTransaction() ([]Transaction, error)
	GetDataTransactionByIDUser(input GetTransactionUserInput)([]Transaction, error)
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) GetAllTransaction() ([]Transaction, error) {
	transactions , err := s.repository.GetTransaction()
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetDataTransactionByIDUser(input GetTransactionUserInput) ([]Transaction, error) {
	transaction , err := s.repository.GetDataTransactionbyIDUser(input.ID)
	if err != nil {
		return transaction,err
	}
	return transaction,nil
}
