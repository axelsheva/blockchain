package services

type IPool interface {
	Lock()
	Unlock()
	Push(*interface{})
	Has(*interface{}) bool
}

type Pool struct {
	locked bool
}
