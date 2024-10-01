package main

import (
	"context"
	GRPCmicro_stc "helloworld/GRPCmicro/stc"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "localhost:3333"
	name = "i am in your walls"
)

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	ct := GRPCmicro_stc.NewTeacherClient(conn)
	cs := GRPCmicro_stc.NewStudentClient(conn)
	cc := GRPCmicro_stc.NewCourseClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r1, err := ct.SayTeacher(ctx, &GRPCmicro_stc.TeacherRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	r2, err := cs.SayStudent(ctx, &GRPCmicro_stc.StudentRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	r3, err := cc.SayCourse(ctx, &GRPCmicro_stc.CourseRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("GreetingT microservice: %s", r1.GetMessage())
	log.Printf("GreetingS microservice: %s", r2.GetMessage())
	log.Printf("GreetingC microservice: %s", r3.GetMessage())

}
