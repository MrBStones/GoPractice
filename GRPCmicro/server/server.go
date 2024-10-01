package main

import (
	"context"
	GRPCmicro_stc "helloworld/GRPCmicro/stc"
	"log"
	"net"

	"google.golang.org/grpc"
)

// should implement the interface myPkgName.InvoicerServer
type TeacherGRPCServer struct {
	// type embedded to comply with Google lib
	GRPCmicro_stc.UnimplementedTeacherServer
}

func (m *TeacherGRPCServer) SayTeacher(ctx context.Context, request *GRPCmicro_stc.TeacherRequest) (*GRPCmicro_stc.TeacherReply, error) {
	log.Println("SayTeacher called microservice")
	return &GRPCmicro_stc.TeacherReply{Message: "hiii i am teach you"}, nil
}

type StudentGRPCServer struct {
	// type embedded to comply with Google lib
	GRPCmicro_stc.UnimplementedStudentServer
}

func (m *StudentGRPCServer) SayStudent(ctx context.Context, request *GRPCmicro_stc.StudentRequest) (*GRPCmicro_stc.StudentReply, error) {
	log.Println("SayStudent called microservice")
	return &GRPCmicro_stc.StudentReply{Message: "ughh why are classes at 8 :("}, nil
}

type CourseGRPCServer struct {
	// type embedded to comply with Google lib
	GRPCmicro_stc.UnimplementedCourseServer
}

func (m *CourseGRPCServer) SayCourse(ctx context.Context, request *GRPCmicro_stc.CourseRequest) (*GRPCmicro_stc.CourseReply, error) {
	log.Println("SayCourse called microservice")
	return &GRPCmicro_stc.CourseReply{Message: "This course we will meet at 8"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	myTeacherServer := &TeacherGRPCServer{}
	GRPCmicro_stc.RegisterTeacherServer(s, myTeacherServer)
	myStudentServer := &StudentGRPCServer{}
	GRPCmicro_stc.RegisterStudentServer(s, myStudentServer)
	myCourseServer := &CourseGRPCServer{}
	GRPCmicro_stc.RegisterCourseServer(s, myCourseServer)

	log.Printf("servers listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
