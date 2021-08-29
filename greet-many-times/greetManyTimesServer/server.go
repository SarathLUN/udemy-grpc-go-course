package main

import (
	"context"
	"github.com/SarathLUN/udemy-grpc-go-course/greet-many-times/greetManyTimesPB"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct {
	greetManyTimesPb.UnimplementedGreetServiceServer
}

func (s server) DoGreet(ctx context.Context, in *greetManyTimesPb.GreetRequest) (*greetManyTimesPb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	firstName := in.GetGreeting().GetFirstName()
	result := "Hello, " + firstName
	res := &greetManyTimesPb.GreetResponse{Result: result}
	return res, nil
}

func (s server) DoGreetManyTimes(req *greetManyTimesPb.GreetManyTimesRequest, stream greetManyTimesPb.GreetService_DoGreetManyTimesServer) error {
	log.Printf("DoGreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello, " + firstName + ", #" + strconv.Itoa(i)
		res := &greetManyTimesPb.GreetManyTimesResponse{Result: result}
		err := stream.Send(res)
		if err != nil {
			log.Fatalf("Error send response: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	log.Println("Starting GreetManyTimes Server ...")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listent: %v", err)
	}
	log.Println("listening on port: 50051")
	s := grpc.NewServer()
	greetManyTimesPb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
