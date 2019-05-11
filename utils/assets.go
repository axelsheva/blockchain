package utils

import (
	"encoding/json"
	"errors"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
	"bitbucket.org/axelsheva/blockchain/models/types"
)

func UnmarshalAsset(bytes []byte, t types.Transaction) (models.IAsset, error) {
	switch t {
	case types.ReferralType:
		var asset assets.Referral
		err := json.Unmarshal(bytes, &asset)
		return &asset, err
	case types.SendType:
		var asset assets.Send
		err := json.Unmarshal(bytes, &asset)
		return &asset, err
	case types.DelegateType:
		var asset assets.Delegate
		err := json.Unmarshal(bytes, &asset)
		return &asset, err
	}
	return nil, errors.New("Invalid transaction type")
}
