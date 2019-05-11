package models

import (
	"encoding/json"
	"errors"
	"time"

	"bitbucket.org/axelsheva/blockchain/models/types"
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
	json.Unmarshaler
	GetAmount() int64
	CalculateFee() int64
}

type ITransaction interface {
	IVerifier
	IApplier
}

type Transaction struct {
	ID              string            `json:"id"`
	BlockID         string            `json:"blockId"`
	Type            types.Transaction `json:"type"`
	SenderPublicKey PublicKey         `json:"senderPublicKey"`
	Fee             int64             `json:"fee"`
	Signature       string            `json:"signature"`
	SecondSignature string            `json:"secondSignature"`
	CreatedAt       time.Time         `json:"createdAt"`
	Salt            string            `json:"salt"`
	Asset           IAsset            `json:"asset"`
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

// func (transaction *Transaction) UnmarshalJSON(data []byte) error {
// 	var tmp struct {
// 		Type types.Transaction `json:"type"`
// 	}

// 	err := json.Unmarshal(data, &tmp)
// 	if err != nil {
// 		return err
// 	}

// 	var asset IAsset
// 	switch tmp.Type {
// 	case types.SendType:
// 		asset = &SendAsset{}
// 	}

// 	transaction.Asset = asset

// 	return json.Unmarshal(data, transaction)
// 	// return nil
// }
