package main

import (
	"time"

	"bitbucket.org/axelsheva/blockchain/configs"
	"bitbucket.org/axelsheva/blockchain/configs/development"
	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/services"
)

func main() {
	configs.Core, configs.Const = configs.Init()

	services.PriorityWorkQueue = models.NewPriorityWorkQueue(1000, 1000)

	err := services.Block.ApplyGenesisBlock(&development.GenesisBlock)
	if err != nil {
		panic(err)
	}

	services.Delegate = services.NewDelegate(configs.Core.ForgeSecret, configs.Const.ActiveDelegatesCount)
	services.Round = services.NewRound(configs.Const.SlotInterval, configs.Const.ActiveDelegatesCount)

	// trustedPeers := strings.Split(os.Getenv("TRUSTED_PEERS"), ",")
	// if len(trustedPeers) == 0 {
	// 	panic("Missing 'TRUSTED_PEERS' environment")
	// }

	// for _, trustedPeer := range trustedPeers {
	// 	go services.StartRPCClient(trustedPeer)
	// }
	// services.StartRPCServer(fmt.Sprint("0.0.0.0:", configs.Core.RPCPort))

	services.PriorityWorkQueue.Push(func() {
		services.Round.Generate(time.Now())
	})

	select {}
}
