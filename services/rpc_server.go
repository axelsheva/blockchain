package services

import (
	"context"
	"log"
	"net"
	"runtime"
	"time"

	"bitbucket.org/axelsheva/blockchain/repositories"

	"bitbucket.org/axelsheva/blockchain/configs"
	"bitbucket.org/axelsheva/blockchain/models/rpc"
	"google.golang.org/grpc"
)

type RPCServer struct{}

func (s *RPCServer) GetSystemInfo(ctx context.Context, in *rpc.SystemInfoRequest) (*rpc.SystemInfoReply, error) {
	log.Printf("[RPC][Server] GetSystemInfo")

	lastBlock := repositories.Blocks.GetLast()

	return &rpc.SystemInfoReply{
		IP:        configs.Core.PublicHost,
		Port:      configs.Core.RPCPort,
		Heigth:    lastBlock.Height,
		OS:        runtime.GOOS,
		Version:   configs.Core.Version,
		Clock:     time.Now().Format(time.RFC3339),
		Broadhash: lastBlock.ID,
	}, nil
}

func StartRPCServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	log.Printf("[RPC][Server] Listening on %s", address)

	rpc.RegisterSystemInfoServer(s, &RPCServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
