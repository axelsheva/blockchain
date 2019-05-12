package main

import (
	"fmt"
	"os"
	"strings"

	"bitbucket.org/axelsheva/blockchain/configs"
	"bitbucket.org/axelsheva/blockchain/configs/development"
	"bitbucket.org/axelsheva/blockchain/services"
)

func main() {
	err := services.Block.ApplyGenesisBlock(&development.GenesisBlock)
	if err != nil {
		panic(err)
	}

	trustedPeers := strings.Split(os.Getenv("TRUSTED_PEERS"), ",")
	if len(trustedPeers) == 0 {
		panic("Missing 'TRUSTED_PEERS' environment")
	}

	for _, trustedPeer := range trustedPeers {
		go services.StartRPCClient(trustedPeer)
	}
	services.StartRPCServer(fmt.Sprint("0.0.0.0:", configs.Core.RPCPort))

	println("Success!")
}
