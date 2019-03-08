package repositories

type IRepository interface {
	AddOne(interface{}) interface{}
	RemoveByID(interface{})
}
