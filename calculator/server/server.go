package main

import (
	"context"
	calculate "github.com/SarathLUN/udemy-grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	calculate.UnimplementedSumServiceServer
}

func (s server) DoSum(ctx context.Context, in *calculate.SumRequest) (*calculate.SumResponse, error) {
	n1 := in.GetInput().GetNum1()
	n2 := in.GetInput().GetNum2()
	r := n1 + n2
	res := &calculate.SumResponse{Result: r}
	return res, nil
}

func main() {
	log.Println("starting server on port 50051...")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen port 50051: %v", err)
	}
	log.Println("server is running on 50051")
	s := grpc.NewServer()
	calculate.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
