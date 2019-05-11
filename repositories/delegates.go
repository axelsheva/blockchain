package repositories

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

var (
	Delegates IDelegatesRepository
)

func init() {
	Delegates = &DelegatesRepository{
		store: make(map[models.PublicKey]*models.Account),
	}
}

type IDelegatesRepository interface {
	Add(account *models.Account)
	Remove(publicKey models.PublicKey)
	GetAll() []*models.Account
}

type DelegatesRepository struct {
	store map[models.PublicKey]*models.Account
}

func (r *DelegatesRepository) Add(account *models.Account) {
	r.store[account.PublicKey] = account
}

func (r *DelegatesRepository) Remove(publicKey models.PublicKey) {
	delete(r.store, publicKey)
}

func (r *DelegatesRepository) GetAll() []*models.Account {
	v := make([]*models.Account, 0, len(r.store))
	for _, value := range r.store {
		v = append(v, value)
	}
	return v
}
