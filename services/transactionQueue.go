package services

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

type ITransactionQueue interface {
	IPool
}

type TransactionQueue struct {
	Pool

	transactions []*models.Transaction
	conflicted   []*models.Transaction
}

func (tq *TransactionQueue) Lock() {
	tq.locked = true
}

func (tq *TransactionQueue) Unlock() {
	tq.locked = false
}

func (tq *TransactionQueue) Push(tx *models.Transaction) {

}
