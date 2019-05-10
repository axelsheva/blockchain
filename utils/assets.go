package utils

import (
	"encoding/json"
	"errors"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
)

func UnmarshalAsset(bytes []byte, t models.TransactionType) (models.IAsset, error) {
	switch t {
	case models.ReferralType:
		var asset assets.Referral
		err := json.Unmarshal(bytes, &asset)
		return &asset, err
	case models.SendType:
		var asset assets.Send
		err := json.Unmarshal(bytes, &asset)
		return &asset, err
	case models.DelegateType:
		var asset assets.Delegate
		err := json.Unmarshal(bytes, &asset)
		return &asset, err
	}
	return nil, errors.New("Invalid transaction type")
}
