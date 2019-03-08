package services

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

type ITransactionPool interface {
	IPool
}

type TransactionPool struct {
	Pool

	transactions map[string]*models.Transaction
	bySender     map[models.PublicKey]*models.Transaction
	byRecipient  map[models.PublicKey]*models.Transaction
}

func (tp *TransactionPool) Lock() {
	tp.locked = true
}

func (tp *TransactionPool) Unlock() {
	tp.locked = false
}

func (tp *TransactionPool) Push(tx *models.Transaction) {

}
