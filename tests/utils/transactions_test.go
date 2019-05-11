package utils_test

import (
	"reflect"
	"testing"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
	"bitbucket.org/axelsheva/blockchain/models/types"
	"bitbucket.org/axelsheva/blockchain/utils"
)

func TestCreateSendTransaction(t *testing.T) {
	transaction := utils.CreateTransaction(models.Transaction{
		Asset: &assets.Send{
			Amount:             100000000,
			RecipientPublicKey: "2",
		},
	}, nil, nil)

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

func TestSendTransactionUnmarshalJSON(t *testing.T) {
	bytes := []byte(`{"id":"c7d80bf1bb220e62735bd388549a87c0cd93b8be30a1ae2f7291ce20d2a94b79","blockId":"cbb9449abb9672d33fa2eb200b1c8b03db7c6572dfb6e59dc334c0ab82b63ab0","type":10,"createdAt":"2019-01-01T00:00:00.000Z","senderPublicKey":"49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a","signature":"226ed984bf3d82b7c332ce48bc976fcc35930d22cb068b2e9de993a4fb3e402d4bdb7077d0923b8dd2c205e6a2473884752615c0787967b218143eec5df1390c","fee":10,"salt":"a7fdae234eeb416e31f5f02571f54a0c","asset":{"recipientPublicKey":"49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a","amount":4500000000000000}}`)

	transaction := models.Transaction{}
	err := utils.UnmarshalTransaction(bytes, &transaction)
	if err != nil {
		panic(err)
	}

	createdAt, err := time.Parse(time.RFC3339, "2019-01-01T00:00:00.000Z")
	if err != nil {
		panic(err)
	}
	expectedTransaction := models.Transaction{
		ID:              "c7d80bf1bb220e62735bd388549a87c0cd93b8be30a1ae2f7291ce20d2a94b79",
		BlockID:         "cbb9449abb9672d33fa2eb200b1c8b03db7c6572dfb6e59dc334c0ab82b63ab0",
		Type:            types.SendType,
		CreatedAt:       createdAt,
		Signature:       "226ed984bf3d82b7c332ce48bc976fcc35930d22cb068b2e9de993a4fb3e402d4bdb7077d0923b8dd2c205e6a2473884752615c0787967b218143eec5df1390c",
		SenderPublicKey: "49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a",
		Salt:            "a7fdae234eeb416e31f5f02571f54a0c",
		Fee:             10,
		Asset: &assets.Send{
			Amount:             4500000000000000,
			RecipientPublicKey: "49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a",
		},
	}

	if !reflect.DeepEqual(transaction, expectedTransaction) {
		t.Errorf("Invalid send transaction unmarshal. Expected: %+v, actual: %+v", expectedTransaction, transaction)
	}
}
