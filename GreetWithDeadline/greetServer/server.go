package main

import (
	"context"
	"github.com/SarathLUN/udemy-grpc-go-course/GreetWithDeadline/greetPb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

type server struct {
	greetPb.UnimplementedGreetServiceServer
}

func (s server) GreetWithDeadline(ctx context.Context, req *greetPb.GreetWithDeadlineRequest) (*greetPb.GreetWithDeadlineResponse, error) {
	log.Printf("GreetWithDeadline function was invoked with %v", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// the client canceled the request
			log.Println("the client cancel request.")
			return nil, status.Error(codes.Canceled, "the client canceled the request.")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello, " + firstName
	res := &greetPb.GreetWithDeadlineResponse{
		Result: result,
	}
	return res, nil
}
func main() {
	log.Println("starting GreetWithDeadline server...")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetPb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
