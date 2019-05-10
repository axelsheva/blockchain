package assets

import (
	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

type Send struct {
	Base
	RecipientPublicKey models.PublicKey `json:"recipientPublicKey"`
	Amount             uint64           `json:"amount"`
}

func (asset *Send) VerifyUnconfirmed(sender *models.Account) error {
	return nil
}

func (asset *Send) GetAmount() uint64 {
	return asset.Amount
}

func (asset *Send) CalculateFee() uint64 {
	return asset.Amount / 10000
}

func (asset *Send) ApplyUnconfirmed(sender *models.Account) {
	sender.Balance -= asset.Amount

	recipient := repositories.Accounts.Get(asset.RecipientPublicKey)
	recipient.Balance += asset.Amount
}

func (asset *Send) UndoUnconfirmed(sender *models.Account) {
	sender.Balance += asset.Amount

	recipient := repositories.Accounts.Get(asset.RecipientPublicKey)
	recipient.Balance -= asset.Amount
}
