package main

import (
	"context"
	GRPC_stc "helloworld/GRPC/stc"
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
	c := GRPC_stc.NewUniversityClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r1, err := c.SayTeacher(ctx, &GRPC_stc.TeacherRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	r2, err := c.SayStudent(ctx, &GRPC_stc.StudentRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	r3, err := c.SayCourse(ctx, &GRPC_stc.CourseRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("GreetingT: %s", r1.GetMessage())
	log.Printf("GreetingS: %s", r2.GetMessage())
	log.Printf("GreetingC: %s", r3.GetMessage())

}
