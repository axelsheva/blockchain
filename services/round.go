package services

import (
	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

var (
	RoundService IRountService
)

func init() {
	RoundService = &Round{
		slotInterval: 10,
	}
}

type IRountService interface {
	generateHashList() []models.HashList
	generateSlots(blockID string, delegates []*models.Account)
	Generate()
	GetMySlot() models.Slot
}

type Round struct {
	slotInterval uint8
}

func (r *Round) generateSlots(blockID string, delegates []*models.Account) {
	// firstSlot := utils.CaltulateFirstSlot()
}

func (r *Round) generateHashList() []models.HashList {
	return nil
}

func (r *Round) Generate() {
	lastBlock := repositories.Blocks.GetLast()
	delegates := Delegate.GetActive()

	r.generateSlots(lastBlock.ID, delegates)
}

func (r *Round) GetMySlot() models.Slot {
	return 0
}
