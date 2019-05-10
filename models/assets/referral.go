package assets

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

type Referral struct {
	Base
	Referral models.PublicKey `json:"referral"`
}
