package main

import (
	"context"
	"fmt"
	PhysicalTime_stc "helloworld/PhysicalTime"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

// should implement the interface myPkgName.InvoicerServer
type SyncTimeServer struct {
	// type embedded to comply with Google lib
	PhysicalTime_stc.UnimplementedSyncTimeServer
}

func (m *SyncTimeServer) GetTime(ctx context.Context, request *PhysicalTime_stc.TimeRequest) (*PhysicalTime_stc.TimeReply, error) {
	serverTime := time.Now().UnixMicro()
	log.Println("--- GetTime ---")
	log.Println("Time back: ", request.Time)
	log.Println("Time sent: ", serverTime)
	log.Println("Diff in mcrscnd: ", (serverTime - request.Time))
	log.Println("Diff in seconds: ", (float64(serverTime-request.Time) * float64(1.0/1000000.0)))
	fmt.Println()

	return &PhysicalTime_stc.TimeReply{Time: serverTime}, nil
}

func main() {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = ":3333"
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	mySyncTimeServer := &SyncTimeServer{}
	PhysicalTime_stc.RegisterSyncTimeServer(s, mySyncTimeServer)

	log.Printf("servers listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
