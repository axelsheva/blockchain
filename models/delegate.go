package models

type Delegate struct {
	Username     string
	Url          string
	MissedBlocks uint64
	ForgedBlocks uint64
	Votes        uint64
}
