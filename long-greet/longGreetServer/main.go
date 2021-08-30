package main

import (
	"github.com/SarathLUN/udemy-grpc-go-course/long-greet/longGreetPb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
	longGreetPb.UnimplementedGreetServiceServer
}

func (s server) LongGreet(stream longGreetPb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with a streaming request")
	var result string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finish reading the client stream
			return stream.SendAndClose(&longGreetPb.LongGreetResponse{Result: result})
		}
		if err != nil {
			log.Fatalf("failed reading client stream: %v\n", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "!, "
	}
	return nil
}

func main() {
	log.Println("starting LongGreet Server")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listent: %v", err)
	}
	s := grpc.NewServer()
	longGreetPb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failt to serve: %v", err)
	}
}
