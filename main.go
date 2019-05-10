package main

import (
	"bitbucket.org/axelsheva/blockchain/configs/development"
	"bitbucket.org/axelsheva/blockchain/services"
)

func main() {
	err := services.Block.ApplyGenesisBlock(&development.GenesisBlock)
	if err != nil {
		panic(err)
	}

	println("Success!")
}
