package main

import (
	"context"
	PhysicalTime_stc "helloworld/PhysicalTime"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

// should implement the interface myPkgName.InvoicerServer
type SyncTimeServer struct {
	// type embedded to comply with Google lib
	PhysicalTime_stc.UnimplementedSyncTimeServer
}

func (m *SyncTimeServer) GetTime(ctx context.Context, request *PhysicalTime_stc.TimeRequest) (*PhysicalTime_stc.TimeReply, error) {
	log.Println("GetTime called")
	return &PhysicalTime_stc.TimeReply{Time: time.Now().Unix()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
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
