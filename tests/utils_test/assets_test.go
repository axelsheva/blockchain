package utils_test

import (
	"encoding/json"
	"testing"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/utils"
)

func TestSendAssetUnmarshal(t *testing.T) {
	bytes := []byte(`{"recipientPublicKey":"1","amount":2}`)
	asset, err := utils.UnmarshalAsset(bytes, models.SendType)
	if err != nil {
		t.Error(err)
	}

	expectedAmount := int64(2)
	if asset.GetAmount() != expectedAmount {
		t.Errorf("Transaction amount is invalid. Expected: %d, actual: %d", expectedAmount, asset.GetAmount())
	}

	actual, err := json.Marshal(asset)
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != string(actual) {
		t.Error("Invalid send asset unmarshal")
	}
}
