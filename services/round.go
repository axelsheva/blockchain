package services

import "bitbucket.org/axelsheva/blockchain/models"

var (
	RoundService IRountService
)

func init() {
	RoundService = &Round{}
}

type IRountService interface {
	GenerateHashList() []models.HashList
	GenerateRound()
	GetMySlot() models.Slot
}

type Round struct {
}

func (r *Round) GenerateHashList() []models.HashList {
	return nil
}

func (r *Round) GenerateRound() {
	// var lastBlock = Block.GetLastBlock()
}

func (r *Round) GetMySlot() models.Slot {
	return 0
}
