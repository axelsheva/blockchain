package models

import "time"

type PeerState uint8

const (
	BANNED PeerState = iota
	DISCONNECTED
	CONNECTED
)

type Peer struct {
	IP            string
	Port          int32
	State         PeerState
	OS            string
	Version       string
	Clock         time.Time
	Broadhash     string
	Height        uint64
	LastBlocksIds []string
}
