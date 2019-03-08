package transactions

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

type Referral struct {
	models.Transaction
	referral models.PublicKey
}
