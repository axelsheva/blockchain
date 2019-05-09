package models

type Message struct {
	ID   string
	Body interface{}
}

type RPC interface {
	Emit(peer Peer, message Message)
	Broadcast(message Message)
}
