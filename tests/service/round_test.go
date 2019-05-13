package service_test

import (
	"reflect"
	"testing"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/services"
)

func TestGenerateHashList(t *testing.T) {
	blockID := "cbb9449abb9672d33fa2eb200b1c8b03db7c6572dfb6e59dc334c0ab82b63ab0"
	delegates := []*models.Account{
		{
			PublicKey: "4fa2e2b4c584f1b914d06e49d3c0d0ff5298fad181f8376ae121ad88f5b8f238",
		},
		{
			PublicKey: "053097685206f4a87a69d8259846012aac3931d9d11b3a06406ecb631e83b376",
		},
		{
			PublicKey: "b7856d4652565e3d32fbb4938f830200f67f9edddd329adaab67cb0a47b896e9",
		},
	}

	roundService := services.RoundService{}
	hashList := roundService.GenerateHashList(blockID, delegates)

	expectedHashList := []*models.HashList{
		{
			PublicKey: "4fa2e2b4c584f1b914d06e49d3c0d0ff5298fad181f8376ae121ad88f5b8f238",
			Hash:      "04029769372b9e8914d44eabd54b0175",
		},
		{
			PublicKey: "053097685206f4a87a69d8259846012aac3931d9d11b3a06406ecb631e83b376",
			Hash:      "5b66f09b456ae8a7474945147aa2792e",
		},
		{
			PublicKey: "b7856d4652565e3d32fbb4938f830200f67f9edddd329adaab67cb0a47b896e9",
			Hash:      "7c87222272e8a7185289ca7e3f8f8d2f",
		},
	}

	if !reflect.DeepEqual(hashList, expectedHashList) {
		t.Errorf("Invalid hash list. Expected: %+v, actual: %+v", expectedHashList, hashList)
	}
}

func TestGenerateSlots(t *testing.T) {
	blockID := "cbb9449abb9672d33fa2eb200b1c8b03db7c6572dfb6e59dc334c0ab82b63ab0"
	delegates := []*models.Account{
		{
			PublicKey: "4fa2e2b4c584f1b914d06e49d3c0d0ff5298fad181f8376ae121ad88f5b8f238",
		},
		{
			PublicKey: "b7856d4652565e3d32fbb4938f830200f67f9edddd329adaab67cb0a47b896e9",
		},
		{
			PublicKey: "053097685206f4a87a69d8259846012aac3931d9d11b3a06406ecb631e83b376",
		},
	}

	roundService := services.RoundService{}
	slots := roundService.GenerateSlots(blockID, delegates, 0)

	expectedSlots := models.Slots{}
	expectedSlots["4fa2e2b4c584f1b914d06e49d3c0d0ff5298fad181f8376ae121ad88f5b8f238"] = 0
	expectedSlots["053097685206f4a87a69d8259846012aac3931d9d11b3a06406ecb631e83b376"] = 1
	expectedSlots["b7856d4652565e3d32fbb4938f830200f67f9edddd329adaab67cb0a47b896e9"] = 2

	if !reflect.DeepEqual(slots, &expectedSlots) {
		t.Errorf("Invalid hash list. Expected: %+v, actual: %+v", expectedSlots, slots)
	}
}

func TestGenerateRound(t *testing.T) {
	// t := time.Parse()

}
