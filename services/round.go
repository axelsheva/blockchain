package services

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"sort"
	"time"

	"bitbucket.org/axelsheva/blockchain/utils"

	"bitbucket.org/axelsheva/blockchain/models"
)

var (
	Round IRountService
)

func init() {
	Round = &RoundService{
		slotInterval:         10,
		activeDelegatesCount: 3,
	}
}

type IRountService interface {
	GenerateHashList(blockID string, delegates []*models.Account) []*models.HashList
	GenerateSlots(blockID string, delegates []*models.Account, firstSlot uint64) *models.Slots
	Generate(t time.Time)
	GetMySlot() models.Slot
}

type RoundService struct {
	slotInterval         uint8
	activeDelegatesCount uint16
}

func (r *RoundService) GenerateSlots(blockID string, delegates []*models.Account, firstSlot uint64) *models.Slots {
	hashList := r.GenerateHashList(blockID, delegates)
	sort.Sort(utils.HashSort(hashList))

	slots := models.Slots{}
	slot := firstSlot
	for _, hash := range hashList {
		slots[hash.PublicKey] = slot
		slot++
	}
	return &slots
}

func (r *RoundService) GenerateHashList(blockID string, delegates []*models.Account) []*models.HashList {
	hashList := make([]*models.HashList, len(delegates))
	for index, delegate := range delegates {
		h := md5.New()
		io.WriteString(h, string(delegate.PublicKey)+blockID)

		hash := hex.EncodeToString(h.Sum(nil))

		hashList[index] = &models.HashList{
			PublicKey: delegate.PublicKey,
			Hash:      hash,
		}
	}
	return hashList
}

func (r *RoundService) Generate(t time.Time) {
	// lastBlock := repositories.Blocks.GetLast()
	// delegates := Delegate.GetActive()
	// firstSlot := utils.CaltulateFirstSlot(t, r.slotInterval, r.activeDelegatesCount)

	// r.generateSlots(lastBlock.ID, delegates, firstSlot)
}

func (r *RoundService) GetMySlot() models.Slot {
	return 0
}
