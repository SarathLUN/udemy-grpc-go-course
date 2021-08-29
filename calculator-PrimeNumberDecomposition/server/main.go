package main

import (
	"context"
	"fmt"
	"github.com/SarathLUN/udemy-grpc-go-course/calculator-PrimeNumberDecomposition/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	calculator.UnimplementedCalculatorServiceServer
}

func (s server) Sum(ctx context.Context, req *calculator.SumRequest) (*calculator.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculator.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func (s server) PrimeNumberDecomposition(req *calculator.PrimeNumberDecompositionRequest, stream calculator.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Printf("Received PrimeNumberDecomposition RPC: %v\n", req)
	number := req.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculator.PrimeNumberDecompositionResponse{PrimeFactor: divisor})
			number = number / divisor
		} else {
			divisor++
			log.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil
}

func main() {
	log.Println("Calculator Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listent: %v", err)
	}
	s := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
