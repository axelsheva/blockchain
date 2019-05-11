package development

import (
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
	"bitbucket.org/axelsheva/blockchain/models/types"
)

// GenesisBlock for developmet
var GenesisBlock models.Block

func init() {
	timestamp, err := time.Parse(time.RFC3339, "2019-01-01T00:00:00.000Z")
	if err != nil {
		panic(err)
	}

	GenesisBlock = models.Block{
		ID:          "cbb9449abb9672d33fa2eb200b1c8b03db7c6572dfb6e59dc334c0ab82b63ab0",
		Version:     1,
		Height:      1,
		PayloadHash: "7e6ba6ec459d96207414f61b67ecd2ecd8c946507bb102a1e47a0ce987e494d0",
		CreatedAt:   timestamp,
		Transactions: []*models.Transaction{
			{
				ID:              "c7d80bf1bb220e62735bd388549a87c0cd93b8be30a1ae2f7291ce20d2a94b79",
				BlockID:         "cbb9449abb9672d33fa2eb200b1c8b03db7c6572dfb6e59dc334c0ab82b63ab0",
				Type:            types.SendType,
				CreatedAt:       timestamp,
				SenderPublicKey: "49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a",
				Salt:            "a7fdae234eeb416e31f5f02571f54a0c",
				Asset: &assets.Send{
					RecipientPublicKey: "95d280b905d41b6bfea2a9f115a78074e20bce67524a6ff1f06d94a9a4d438ea",
					Amount:             4500000000000000,
				},
			},
			{
				ID:              "06b79b6bef1234318893708898e7b61ada5671d171b2d8a0737492071fce3574",
				BlockID:         "cbb9449abb9672d33fa2eb200b1c8b03db7c6572dfb6e59dc334c0ab82b63ab0",
				Type:            types.DelegateType,
				CreatedAt:       timestamp,
				SenderPublicKey: "4fa2e2b4c584f1b914d06e49d3c0d0ff5298fad181f8376ae121ad88f5b8f238",
				Salt:            "c41398ec44dbdec4ccb16d372e79aa0a",
				Asset: &assets.Delegate{
					Username: "Delegate 1",
				},
			},
			{
				ID:              "4e87f1527c193a55c0661e2d91b58d1167bedf9211da6a765081319333c30b3c",
				BlockID:         "cbb9449abb9672d33fa2eb200b1c8b03db7c6572dfb6e59dc334c0ab82b63ab0",
				Type:            types.DelegateType,
				CreatedAt:       timestamp,
				SenderPublicKey: "053097685206f4a87a69d8259846012aac3931d9d11b3a06406ecb631e83b376",
				Salt:            "b8943007c493fed26a60f68128378cf9",
				Asset: &assets.Delegate{
					Username: "Delegate 2",
				},
			},
			{
				ID:              "a543f11a6af530944e447f191a9e424050981c72a0194f985f4b61118820df4d",
				BlockID:         "cbb9449abb9672d33fa2eb200b1c8b03db7c6572dfb6e59dc334c0ab82b63ab0",
				Type:            types.DelegateType,
				CreatedAt:       timestamp,
				SenderPublicKey: "b7856d4652565e3d32fbb4938f830200f67f9edddd329adaab67cb0a47b896e9",
				Salt:            "e014fa1c15814d7f2a590193063f631b",
				Asset: &assets.Delegate{
					Username: "Delegate 3",
				},
			},
		},
	}
}
