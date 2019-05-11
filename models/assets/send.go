package assets

import (
	"encoding/json"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

type Send struct {
	Base
	RecipientPublicKey models.PublicKey `json:"recipientPublicKey"`
	Amount             int64            `json:"amount"`
}

func (asset *Send) VerifyUnconfirmed(sender *models.Account) error {
	return nil
}

func (asset *Send) GetAmount() int64 {
	return asset.Amount
}

func (asset *Send) CalculateFee() int64 {
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

func (asset *Send) UnmarshalJSON(data []byte) error {
	var tmp struct {
		RecipientPublicKey string `json:"recipientPublicKey"`
		Amount             int64  `json:"amount"`
	}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	asset.Amount = tmp.Amount
	asset.RecipientPublicKey = models.PublicKey(tmp.RecipientPublicKey)

	return nil
}
