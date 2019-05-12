package services

import (
	"context"
	"log"
	"time"

	"bitbucket.org/axelsheva/blockchain/models/rpc"
	"google.golang.org/grpc"
)

func StartRPCClient(address string) {
	time.Sleep(5000 * time.Millisecond)
	log.Println("Connecting to", address)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc.NewSystemInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetSystemInfo(ctx, &rpc.SystemInfoRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("From server: %+v", r)
}
