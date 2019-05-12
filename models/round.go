package models

type HashList struct {
	Hash      string
	PublicKey PublicKey
}

type Round struct {
	ID          uint64
	Slots       map[PublicKey]Slot
	StartHeight uint64
	EndHeight   uint64
}
