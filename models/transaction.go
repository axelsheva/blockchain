package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type TransactionType uint8

const (
	ReferralType  TransactionType = 0
	SendType                      = 10
	SignatureType                 = 20
	DelegateType                  = 30
)

const SALT_LENGTH = 32

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
	GetAmount() int64
	CalculateFee() int64
}

// func (asset IAsset) UnmarshalJSON(data []byte) error {
// 	println("[asset][UnmarshalJSON]")

// 	return nil
// }

type ITransaction interface {
	IVerifier
	IApplier
}

type Transaction struct {
	ID              string          `json:"id"`
	BlockID         string          `json:"blockId"`
	Type            TransactionType `json:"type"`
	SenderPublicKey PublicKey       `json:"senderPublicKey"`
	Fee             int64          `json:"fee"`
	Signature       string          `json:"signature"`
	SecondSignature string          `json:"secondSignature"`
	CreatedAt       time.Time       `json:"createdAt"`
	Salt            string          `json:"salt"`
	Asset           IAsset          `json:"asset"`
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

func (transaction *Transaction) UnmarshalJSON(data []byte) error {
	json.Unmarshal(data, *transaction)

	// var s string
	// err := json.Unmarshal(data, s)
	// if err != nil {
	// 	return err
	// }
	fmt.Printf("s: %+v\n\n", transaction)

	// utils.UnmarshalAsset(data)

	// return json.Unmarshal(data, transaction.Asset)
	return nil
}
