package utils_test

import (
	"encoding/hex"
	"testing"

	"bitbucket.org/axelsheva/blockchain/utils"
)

func TestGenerateKeyPair(t *testing.T) {
	secret := "endorse bicycle chunk simple improve paper use radar crazy rain bullet short"

	keyPair := utils.GenerateKeyPair(secret)

	expectedPublicKey := "af682caed24354ad611164b95fa6c43b63c5ebc88712e2cd1f1088a9ba0e167f"
	actualPublicKey := hex.EncodeToString(keyPair.PublicKey.Bytes)
	if actualPublicKey != expectedPublicKey {
		t.Errorf("Invalid public key. Expected: %s, actual: %s", expectedPublicKey, actualPublicKey)
	}

	expectedPrivateKey := "96aafe4f69a36197251d181df0d88f255a3abcb8046731b6036ed75c4eca65c4af682caed24354ad611164b95fa6c43b63c5ebc88712e2cd1f1088a9ba0e167f"
	actualPrivateKey := hex.EncodeToString(keyPair.SecretKey.Bytes)
	if actualPrivateKey != expectedPrivateKey {
		t.Errorf("Invalid private key. Expected: %s, actual: %s", expectedPrivateKey, actualPrivateKey)
	}
}
