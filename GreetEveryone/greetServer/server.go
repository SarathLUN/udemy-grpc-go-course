package main

import (
	"github.com/SarathLUN/udemy-grpc-go-course/GreetEveryone/greetPb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
	greetPb.UnimplementedGreetServiceServer
}

func (s server) GreetEveryone(stream greetPb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone is invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("failed to read client streaming: %v", err)
			return err
		}

		firstName := req.GetGreeting().GetFirstName()
		result := "Hello, " + firstName + "! "
		err = stream.Send(&greetPb.GreetEveryoneResponse{Result: result})
		if err != nil {
			log.Fatalf("error while sending data to client")
			return err
		}
	}
}

func main() {
	log.Println("starting GreetEveryone server...")
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
