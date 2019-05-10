package assets

import (
	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

type Referral struct {
	Base
	Referral models.PublicKey `json:"referral"`
}

func (asset *Referral) VerifyUnconfirmed(sender *models.Account) error {
	return nil
}

func (asset *Referral) ApplyUnconfirmed(sender *models.Account) {
	referral := repositories.Accounts.Get(asset.Referral)

	sender.Referral = referral
}

func (asset *Referral) UndoUnconfirmed(sender *models.Account) {
	sender.Referral = nil
}
