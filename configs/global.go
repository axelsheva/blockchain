package configs

import (
	"os"
	"strconv"

	"bitbucket.org/axelsheva/blockchain/configs/development"
	"bitbucket.org/axelsheva/blockchain/models"
)

var Core *models.CoreConfig
var Const models.Const

func Init() (*models.CoreConfig, models.Const) {
	RPCPort, err := strconv.ParseInt(os.Getenv("RPC_PORT"), 10, 32)
	if err != nil {
		panic("Missing or invalid 'RPC_PORT' environment")
	}
	publicHost := os.Getenv("PUBLIC_HOST")
	if publicHost == "" {
		panic("Missing 'PUBLIC_HOST' environment")
	}
	forgeSecret := os.Getenv("FORGE_SECRET")
	if forgeSecret == "" {
		panic("Missing 'FORGE_SECRET' environment")
	}

	coreConfig := models.CoreConfig{
		RPCPort:     int32(RPCPort),
		PublicHost:  publicHost,
		Version:     "0.0.1",
		ForgeSecret: forgeSecret,
	}
	constConfig := development.Const

	return &coreConfig, constConfig
}
