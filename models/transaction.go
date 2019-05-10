package models

import (
	"errors"
	"time"
)

type TransactionType uint8

const (
	Referral TransactionType = iota
	Send
)

type IVerifier interface {
	VerifyUnconfirmed(sender *Account) error
}

type IApplier interface {
	ApplyUnconfirmed(sender *Account)
	UndoUnconfirmed(sender *Account)
}

type IAsset interface {
	IVerifier
	IApplier
	GetAmount() uint64
	CalculateFee() uint64
}

type ITransaction interface {
	IVerifier
	IApplier
}

type Transaction struct {
	ID              string          `json:"id"`
	BlockID         string          `json:"block_id"`
	Type            TransactionType `json:"type"`
	SenderPublicKey PublicKey       `json:"sender_public_key"`
	Fee             uint64          `json:"fee"`
	Signature       string          `json:"signature"`
	SecondSignature string          `json:"second_signature"`
	CreatedAt       time.Time       `json:"created_at"`
	Salt            string          `json:"salt"`
	Asset           IAsset          `json:"asset"`
}

func CreateTransaction(data Transaction) *Transaction {
	transaction := Transaction{
		Fee:       data.Asset.CalculateFee(),
		Asset:     data.Asset,
		CreatedAt: time.Now(),
	}

	return &transaction
}

func (transaction *Transaction) VerifyUnconfirmed(sender *Account) error {
	amount := transaction.Fee + transaction.Asset.GetAmount()
	if sender.Balance < amount {
		return errors.New("Not enough money")
	}
	return transaction.Asset.VerifyUnconfirmed(sender)
}

func (transaction *Transaction) ApplyUnconfirmed(sender *Account) {
	sender.Balance -= transaction.Fee

	transaction.Asset.ApplyUnconfirmed(sender)
}

func (transaction *Transaction) UndoUnconfirmed(sender *Account) {
	sender.Balance += transaction.Fee

	transaction.Asset.UndoUnconfirmed(sender)
}
