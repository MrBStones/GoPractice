package main

import (
	"context"
	GRPC_stc "helloworld/GRPC/stc"
	"log"
	"net"

	"google.golang.org/grpc"
)

// should implement the interface myPkgName.InvoicerServer
type myGRPCServer struct {
	// type embedded to comply with Google lib
	GRPC_stc.UnimplementedUniversityServer
}

func (m *myGRPCServer) SayTeacher(ctx context.Context, request *GRPC_stc.TeacherRequest) (*GRPC_stc.TeacherReply, error) {
	log.Println("SayTeacher called")
	return &GRPC_stc.TeacherReply{Message: "hiii i am teach you"}, nil
}

func (m *myGRPCServer) SayStudent(ctx context.Context, request *GRPC_stc.StudentRequest) (*GRPC_stc.StudentReply, error) {
	log.Println("SayStudent called")
	return &GRPC_stc.StudentReply{Message: "ughh why are classes at 8 :("}, nil
}

func (m *myGRPCServer) SayCourse(ctx context.Context, request *GRPC_stc.CourseRequest) (*GRPC_stc.CourseReply, error) {
	log.Println("SayCourse called")
	return &GRPC_stc.CourseReply{Message: "This course we will meet at 8"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	myUniversityServer := &myGRPCServer{}
	GRPC_stc.RegisterUniversityServer(s, myUniversityServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
