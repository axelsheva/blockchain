package models

import "time"

type Block struct {
	ID                 string         `json:"id"`
	PreviousBlockID    string         `json:"previousBlockId"`
	PayloadHash        string         `json:"payloadHash"`
	Signature          string         `json:"signature"`
	GeneratorPublicKey PublicKey      `json:"generatorPublicKey"`
	Height             uint64         `json:"height"`
	Amount             int64          `json:"amount"`
	Fee                int64          `json:"fee"`
	Transactions       []*Transaction `json:"transactions"`
	Version            uint8          `json:"version"`
	CreatedAt          time.Time      `json:"createdAt"`
}

type IBlockController interface {
	OnReceive(block *Block) error
	Generate() (*Block, error)
}
