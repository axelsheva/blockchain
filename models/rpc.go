package models

type MessageType uint8

const (
	MessageWithBlock MessageType = iota
	MessageWithTransaction
)

type Message struct {
	ID   string
	Type MessageType
	Body interface{}
}

type RPC interface {
	Emit(peer Peer, message Message)
	Broadcast(message Message)
}
