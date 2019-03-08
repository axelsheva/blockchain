package repositories

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

var (
	Delegates IDelegatesRepository
)

func init() {
	Delegates = &DelegatesRepository{
		byPublicKey: make(map[models.PublicKey]*models.Account),
	}
}

type IDelegatesRepository interface {
	AddOne(account models.Account) *models.Account
	RemoveByPublicKey(publicKey models.PublicKey)
}

type DelegatesRepository struct {
	byPublicKey map[models.PublicKey]*models.Account
}

func (dr *DelegatesRepository) AddOne(account models.Account) *models.Account {
	dr.byPublicKey[account.PublicKey] = &account
	return &account
}

func (dr *DelegatesRepository) RemoveByPublicKey(publicKey models.PublicKey) {
	delete(dr.byPublicKey, publicKey)
}
