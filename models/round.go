package models

type HashList struct {
	hash               string
	generatorPublicKey string
}

type Round struct {
	ID          uint64
	Slots       map[PublicKey]Slot
	StartHeight uint64
	EndHeight   uint64
}
