package utils

import (
	"encoding/json"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
	"bitbucket.org/axelsheva/blockchain/models/types"
	"github.com/jamesruan/sodium"
)

func CreateTransaction(data models.Transaction, keyPair *sodium.SignKP, secondKeyPair *sodium.SignKP) *models.Transaction {
	transaction := models.Transaction{
		Fee:       data.Asset.CalculateFee(),
		Asset:     data.Asset,
		Salt:      RandStringBytesMask(models.SALT_LENGTH),
		CreatedAt: time.Now(),
	}

	return &transaction
}

func UnmarshalTransaction(data []byte, transaction *models.Transaction) error {
	var tmp struct {
		Type types.Transaction `json:"type"`
	}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	var asset models.IAsset
	switch tmp.Type {
	case types.SendType:
		asset = &assets.Send{}
	}

	transaction.Asset = asset

	return json.Unmarshal(data, transaction)
}
