package models

type PeerState uint8

const (
	BANNED PeerState = iota
	DISCONNECTED
	CONNECTED
)

type Peer struct {
	IP            string
	Port          uint16
	State         PeerState
	OS            string
	Version       string
	Clock         uint64
	Broadhash     string
	Height        uint64
	Socket        string
	LastBlocksIds []string
}
