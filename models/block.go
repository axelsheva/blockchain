package models

type Block struct {
	ID                 string        `json:"id"`
	PreviousBlockID    string        `json:"previous_block_id"`
	PayloadHash        string        `json:"payload_hash"`
	Signature          string        `json:"signature"`
	GeneratorPublicKey PublicKey     `json:"generator_public_key"`
	Height             uint64        `json:"height"`
	Amount             uint64        `json:"amount"`
	Fee                uint64        `json:"fee"`
	Transactions       []Transaction `json:"transactions"`
	Version            uint8         `json:"version"`
}

type IBlockController interface {
	OnReceive(block *Block) error
	Generate() (*Block, error)
}
