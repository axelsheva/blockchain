package utils_test

import (
	"testing"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
	"bitbucket.org/axelsheva/blockchain/utils"
)

func TestCreateSendTransaction(t *testing.T) {
	transaction := utils.CreateTransaction(models.Transaction{
		Asset: &assets.Send{
			Amount:             100000000,
			RecipientPublicKey: "2",
		},
	}, models.KeyPair{
		PrivateKey: "",
		PublicKey:  "",
	}, models.KeyPair{
		PrivateKey: "",
		PublicKey:  "",
	})

	expectedFee := int64(10000)
	if transaction.Fee != expectedFee {
		t.Errorf("Transaction fee is invalid. Expected: %d, actual: %d", expectedFee, transaction.Fee)
	}

	if transaction.Salt == "" {
		t.Errorf("Transaction salt is missing")
	}

	if len(transaction.Salt) != models.SALT_LENGTH {
		t.Errorf("Transaction salt has invalid length. Expected: %d, actual: %d", models.SALT_LENGTH, len(transaction.Salt))
	}
}
