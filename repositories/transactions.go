package repositories

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

var (
	Transactions ITransactionsRepository
)

func init() {
	Transactions = &TransactionsRepository{
		transactionsByID: make(map[string]*models.Transaction),
	}
}

type ITransactionsRepository interface {
	AddOne(tx models.Transaction) *models.Transaction
	RemoveByID(ID string)
}

type TransactionsRepository struct {
	transactionsByID map[string]*models.Transaction
}

func (tr *TransactionsRepository) AddOne(tx models.Transaction) *models.Transaction {
	tr.transactionsByID[tx.ID] = &tx
	return &tx
}

func (tr *TransactionsRepository) RemoveByID(ID string) {
	delete(tr.transactionsByID, ID)
}
