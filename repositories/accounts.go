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
	Add(account *models.Account) *models.Account
	Get(publicKey models.PublicKey) *models.Account
}

type AccountsRepository struct {
	byPublicKey map[models.PublicKey]*models.Account
}

func (repository *AccountsRepository) Add(account *models.Account) *models.Account {
	repository.byPublicKey[account.PublicKey] = account
	return account
}

func (repository *AccountsRepository) Get(publicKey models.PublicKey) *models.Account {
	return repository.byPublicKey[publicKey]
}
