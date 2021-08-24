package main

import (
	"context"
	"github.com/SarathLUN/udemy-grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s server) DoGreet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	firstName := in.GetGreeting().GetFirstName()
	result := "Hello, " + firstName
	res := &greetpb.GreetResponse{Result: result}
	return res, nil
}

func main() {
	log.Println("Hello world!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listent: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
