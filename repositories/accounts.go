package repositories

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

var (
	Accounts IAccountsRepository
)

func init() {
	Accounts = &AccountsRepository{
		byPublicKey: make(map[models.PublicKey]*models.Account),
	}
}

type IAccountsRepository interface {
	AddOne(account models.Account) *models.Account
	RemoveByPublicKey(publicKey models.PublicKey)
}

type AccountsRepository struct {
	byPublicKey map[models.PublicKey]*models.Account
}

func (dr *AccountsRepository) AddOne(account models.Account) *models.Account {
	dr.byPublicKey[account.PublicKey] = &account
	return &account
}

func (dr *AccountsRepository) RemoveByPublicKey(publicKey models.PublicKey) {
	delete(dr.byPublicKey, publicKey)
}
