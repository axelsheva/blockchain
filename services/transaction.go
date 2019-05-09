package services

import (
	"bytes"

	"bitbucket.org/axelsheva/blockchain/models"
)

type ITransaction interface {
	Create(data interface{}) *interface{}

	Validate(data interface{}) error

	Verify(data interface{}) error

	ApplyUnconfirmed(data interface{}) error
	UndoUnconfirmed(data interface{}) error

	GetBytes(data interface{}) bytes.Buffer

	Apply(data interface{}) error
	Undo(data interface{}) error
}

type TransactionService interface {
	ITransaction

	CalculateID(tx *models.Transaction) string
	CalculateHash(tx *models.Transaction) bytes.Buffer
	ToBytes(tx *models.Transaction) bytes.Buffer

	IsConfirmed(tx *models.Transaction) bool
}
