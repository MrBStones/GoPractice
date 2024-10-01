package main

import (
	"context"
	PhysicalTime_stc "helloworld/PhysicalTime"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "localhost:3333"
)

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	ct := PhysicalTime_stc.NewSyncTimeClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sendTime := time.Now().Unix()
	r, err := ct.GetTime(ctx, &PhysicalTime_stc.TimeRequest{Time: sendTime})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println("Time gotten back: ", r.Time)
	log.Println("Diff: ", (r.Time - sendTime))

}
