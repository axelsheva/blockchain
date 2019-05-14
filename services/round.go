package services

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"sort"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
	"bitbucket.org/axelsheva/blockchain/utils"
)

var (
	Round IRountService
)

func NewRound(slotInterval uint8, activeDelegatesCount uint16) IRountService {
	log.Printf("[Service][Round][NewRound] Slot interval: %ds, active delegates count: %d", slotInterval, activeDelegatesCount)

	return &RoundService{
		slotInterval:         slotInterval,
		activeDelegatesCount: activeDelegatesCount,
	}
}

type IRountService interface {
	GenerateHashList(blockID string, delegates []*models.Account) []*models.HashList
	GenerateSlots(blockID string, delegates []*models.Account, firstSlot int64) *models.Slots
	Generate(t time.Time)
	GetMySlot() models.Slot
}

type RoundService struct {
	slotInterval         uint8
	activeDelegatesCount uint16
}

func (r *RoundService) GenerateSlots(blockID string, delegates []*models.Account, firstSlot int64) *models.Slots {
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

func (r *RoundService) Finish(d time.Duration) {
	time.Sleep(d)

	PriorityWorkQueue.Push(func() {
		r.Generate(time.Now())
	})
}

func (r *RoundService) Generate(t time.Time) {
	log.Printf("[Service][Round][Generate] Generate new round for time: %s", t.Format(time.RFC3339))
	lastBlock := repositories.Blocks.GetLast()
	delegates := Delegate.GetActive()
	firstSlot := utils.CalculateFirstSlot(t, r.slotInterval, r.activeDelegatesCount)

	round := models.Round{
		Slots: *r.GenerateSlots(lastBlock.ID, delegates, firstSlot),
	}
	repositories.Rounds.Push(&round)

	Delegate.Forge(&round)

	lastSlot := round.GetLastSlot()

	// TODO: check without throttling
	throttling := int64(100000000)
	finishTime := time.Unix((lastSlot+1)*int64(r.slotInterval), throttling)
	diff := finishTime.Sub(time.Now()).Nanoseconds() / int64(time.Millisecond)

	duration := time.Duration(diff * int64(time.Millisecond))
	log.Printf("[Service][Round][Generate] Round will be finish after %s", duration.String())

	go r.Finish(duration)
}

func (r *RoundService) GetMySlot() models.Slot {
	return 0
}
