package main

import (
	"context"
	"fmt"
	"github.com/SarathLUN/udemy-grpc-go-course/calculator-square-root/calculatorPb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
	"net"
)

type server struct {
	calculatorPb.UnimplementedCalculatorServiceServer
}

func (s server) SquareRoot(_ context.Context, req *calculatorPb.SquareRootRequest) (*calculatorPb.SquareRootResponse, error) {
	number := req.GetNumber()
	log.Printf("Received SquareRoot RPC: %v \n", number)
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Receive a negative number: %v", number),
		)
	}
	return &calculatorPb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	// console message server start
	log.Println("Calculator Server")
	// create listener
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create server
	s := grpc.NewServer()
	calculatorPb.RegisterCalculatorServiceServer(s, &server{})
	// serve
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
