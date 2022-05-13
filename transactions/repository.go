package transactions

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetTransaction() ([]Transaction, error)
	GetDataTransactionbyIDUser(id int) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTransaction() ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) GetDataTransactionbyIDUser(id int) ([]Transaction, error) {
	//TODO implement me
	var transaction []Transaction

	err := r.db.Preload("User").Where("user_id = ?", id).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
