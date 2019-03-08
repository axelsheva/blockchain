package services

import (
	"bitbucket.org/axelsheva/blockchain/models"
	"time"
)

var (
	Slot ISlot
)

func init() {
	Slot = &SlotService{
		interval:        10,
		activeDelegates: 3,
		epochTime:       time.Date(2016, time.December, 1, 17, 0, 0, 0, time.UTC),
	}
}

type ISlot interface {
	BeginEpochTime() time.Time
	GetEpochTime(time *time.Time) models.EpochTime
	GetRealTime(epochTime models.EpochTime) time.Time
	GetSlotNumber(epochTime *models.EpochTime) models.Slot
	GetSlotTime(slot models.Slot) models.EpochTime
	GetNextSlot() models.Slot
}

type SlotService struct {
	epochTime       time.Time
	interval        int64
	activeDelegates uint8
}

func (ss *SlotService) BeginEpochTime() time.Time {
	return ss.epochTime
}

func (ss *SlotService) GetEpochTime(t *time.Time) models.EpochTime {
	if t == nil {
		var tmp = time.Now()
		t = &tmp
	}

	return t.Unix() - ss.epochTime.Unix()
}

func (ss *SlotService) GetRealTime(epochTime models.EpochTime) time.Time {
	return time.Unix(ss.epochTime.Unix()+epochTime, 0)
}

func (ss *SlotService) GetSlotNumber(epochTime *models.EpochTime) models.Slot {
	if epochTime == nil {
		var tmp = ss.GetEpochTime(nil)
		epochTime = &tmp
	}

	return ss.epochTime.Unix() / ss.interval
}

func (ss *SlotService) GetSlotTime(slot models.Slot) models.EpochTime {
	return slot * ss.interval
}

func (ss *SlotService) GetNextSlot() models.Slot {
	return ss.GetSlotNumber(nil) + 1
}
