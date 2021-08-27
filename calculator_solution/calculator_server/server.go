package main

import (
	"context"
	"github.com/SarathLUN/udemy-grpc-go-course/calculator_solution/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	calculatorpb.UnimplementedDoSumServer
}

func (s server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Server got request(type = %T), value = %v", req, req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{Result: sum}
	return res, nil
}

func main() {
	log.Println("Calculator Server, port:50051")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error: listener port 50051: %v", err)
	}
	log.Println("Successful: listener started!")
	s := grpc.NewServer()
	calculatorpb.RegisterDoSumServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error: serve: %v", err)
	}

}
