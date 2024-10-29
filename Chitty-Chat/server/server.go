package main

import (
	"context"
	"fmt"
	pc "helloworld/Chitty-Chat"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type LeaderServer struct {
	// type embedded to comply with Google lib
	pc.UnimplementedLeaderServer
	mutex sync.Mutex
	sum   int
}

func (s *LeaderServer) SayLeader(ctx context.Context, request *pc.LeaderRequest) (*pc.LeaderReply, error) {
	s.mutex.Lock()
	log.Println("SayLeader called! Current sum:", s.sum)
	s.sum += int(request.ToAdd)
	s.mutex.Unlock()
	return &pc.LeaderReply{Message: fmt.Sprintf("sum = %d", s.sum)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	myServer := &LeaderServer{}
	pc.RegisterLeaderServer(s, myServer)

	log.Printf("servers listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
