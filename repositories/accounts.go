package repositories

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

var (
	Accounts IAccountsRepository
)

func init() {
	Accounts = &AccountsRepository{
		store: make(map[models.PublicKey]*models.Account),
	}
}

type IAccountsRepository interface {
	Add(account *models.Account)
	Get(publicKey models.PublicKey) *models.Account
	Remove(publicKey models.PublicKey)
}

type AccountsRepository struct {
	store map[models.PublicKey]*models.Account
}

func (r *AccountsRepository) Add(account *models.Account) {
	r.store[account.PublicKey] = account
}

func (r *AccountsRepository) Get(publicKey models.PublicKey) *models.Account {
	account := r.store[publicKey]

	if account == nil {
		account = &models.Account{
			PublicKey: publicKey,
		}
	}

	return account
}

func (r *AccountsRepository) Remove(publicKey models.PublicKey) {
	delete(r.store, publicKey)
}
