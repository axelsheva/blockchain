package models

type Delegate struct {
	Username     string `json:"username"`
	MissedBlocks uint64 `json:"missedBlocks"`
	ForgedBlocks uint64 `json:"forgedBlocks"`
	Votes        uint64 `json:"votes"`
}
