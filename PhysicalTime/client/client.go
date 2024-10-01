package main

import (
	"context"
	"fmt"
	PhysicalTime_stc "helloworld/PhysicalTime"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr1 = "192.168.1.157:3333"
	addr2 = "localhost:3334"
)

func main() {
	ts(addr1)
	ts(addr2)
}

func ts(addr string) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	ct := PhysicalTime_stc.NewSyncTimeClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sendTime := time.Now().UnixMicro()
	r, err := ct.GetTime(ctx, &PhysicalTime_stc.TimeRequest{Time: sendTime})
	reciveTime := time.Now().UnixMicro()
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println("Address: ", addr)
	log.Println("Time sent: ", sendTime)
	log.Println("Time back: ", r.Time)
	log.Println("Diff in mcrscnd: ", (r.Time - sendTime))
	log.Println("Diff in seconds: ", (float64(r.Time-sendTime) * float64(1.0/1000000.0)))
	log.Println("Round-trip time in seconds: ", (float64(reciveTime-sendTime) * float64(1.0/1000000.0)))
	fmt.Println()
}
