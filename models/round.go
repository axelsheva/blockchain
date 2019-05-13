package models

type HashList struct {
	Hash      string
	PublicKey PublicKey
}

type Round struct {
	Slots Slots
}

func (round *Round) GetLastSlot() Slot {
	maxSlot := int64(0)
	for _, slot := range round.Slots {
		if slot > maxSlot {
			maxSlot = slot
		}
	}
	return maxSlot
}
