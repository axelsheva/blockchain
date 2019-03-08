package transactions

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

type Send struct {
	models.Transaction
	recepient models.PublicKey
	amount    uint64
}
