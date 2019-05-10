package utils

import (
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
)

func CreateTransaction(data models.Transaction, keyPair models.KeyPair, secondKeyPair models.KeyPair) *models.Transaction {
	transaction := models.Transaction{
		Fee:       data.Asset.CalculateFee(),
		Asset:     data.Asset,
		Salt:      RandStringBytesMask(models.SALT_LENGTH),
		CreatedAt: time.Now(),
	}

	return &transaction
}
