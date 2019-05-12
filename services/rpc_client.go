package services

import (
	"context"
	"log"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"

	"bitbucket.org/axelsheva/blockchain/models/rpc"
	"bitbucket.org/axelsheva/blockchain/repositories"
	"google.golang.org/grpc"
)

func GetSystemInfo(c rpc.SystemInfoClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetSystemInfo(ctx, &rpc.SystemInfoRequest{})
	if err != nil {
		log.Fatalf("[RPC][Client][GetSystemInfo] Could not get system info: %v", err)
	}

	// TODO: migrate this logic
	clock, err := time.Parse(time.RFC3339, r.Clock)
	if err != nil {
		panic(err)
	}

	peer := models.Peer{
		IP:        r.IP,
		Port:      r.Port,
		State:     models.CONNECTED,
		OS:        r.OS,
		Version:   r.Version,
		Clock:     clock,
		Broadhash: r.Broadhash,
		Height:    r.Heigth,
	}
	repositories.Peers.Add(&peer)
}

func StartRPCClient(address string) {
	time.Sleep(5000 * time.Millisecond)
	log.Println("[RPC][Client] Connecting to", address)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[RPC][Client] Did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc.NewSystemInfoClient(conn)

	GetSystemInfo(c)
}
