package main

import (
	"time"

	"bitbucket.org/axelsheva/blockchain/configs"
	"bitbucket.org/axelsheva/blockchain/configs/development"
	"bitbucket.org/axelsheva/blockchain/services"
)

func main() {
	configs.Init()
	services.InitDelegate(configs.Core.ForgeSecret, configs.Const.ActiveDelegatesCount)
	services.InitRound(configs.Const.SlotInterval, configs.Const.ActiveDelegatesCount)

	err := services.Block.ApplyGenesisBlock(&development.GenesisBlock)
	if err != nil {
		panic(err)
	}

	// trustedPeers := strings.Split(os.Getenv("TRUSTED_PEERS"), ",")
	// if len(trustedPeers) == 0 {
	// 	panic("Missing 'TRUSTED_PEERS' environment")
	// }

	// for _, trustedPeer := range trustedPeers {
	// 	go services.StartRPCClient(trustedPeer)
	// }
	// services.StartRPCServer(fmt.Sprint("0.0.0.0:", configs.Core.RPCPort))

	services.Round.Generate(time.Now())
}
