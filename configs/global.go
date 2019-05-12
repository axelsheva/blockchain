package configs

import (
	"os"
	"strconv"

	"bitbucket.org/axelsheva/blockchain/models"
)

var Core models.CoreConfig

func Init() {
	RPCPort, err := strconv.ParseInt(os.Getenv("RPC_PORT"), 10, 32)
	if err != nil {
		panic("Missing or invalid 'RPC_PORT' environment")
	}
	publicHost := os.Getenv("PUBLIC_HOST")
	if publicHost == "" {
		panic("Missing 'PUBLIC_HOST' environment")
	}

	Core = models.CoreConfig{
		RPCPort:    int32(RPCPort),
		PublicHost: publicHost,
		Version:    "0.0.1",
	}
}
