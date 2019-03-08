package models

import "time"

type TransactionType uint8

const (
	Referral TransactionType = iota
	Send
)

type Transaction struct {
	ID              string
	BlockID         string
	Type            TransactionType
	Sender          PublicKey
	Fee             uint64
	Signature       string
	SecondSignature string
	CreatedAt       time.Time
	Salt            string
}
