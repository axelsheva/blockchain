package assets

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

type Delegate struct {
	Base
	Username string `json:"username"`
}

func (*Delegate) CalculateFee() int64 {
	return 1000000000
}

func (asset *Delegate) VerifyUnconfirmed(sender *models.Account) error {
	return nil
}

func (asset *Delegate) ApplyUnconfirmed(sender *models.Account) {
	delegate := models.Delegate{
		Username: asset.Username,
	}

	sender.Delegate = &delegate
}

func (asset *Delegate) UndoUnconfirmed(sender *models.Account) {
	sender.Delegate = nil
}
