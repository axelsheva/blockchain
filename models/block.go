package models

type Block struct {
	ID              string
	PreviousBlockID string
	PayloadHash     string
	Signature       string
	Generator       PublicKey
	Height          uint64
	Amount          uint64
	Fee             uint64
	Transactions    []Transaction
	Version         uint8
}
