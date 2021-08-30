package main

import (
	"github.com/SarathLUN/udemy-grpc-go-course/calculator-compute-average/calculatorPb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
	calculatorPb.UnimplementedCalculatorServiceServer
}

func (s server) ComputeAverage(stream calculatorPb.CalculatorService_ComputeAverageServer) error{
	log.Println("Received ComputeAverage RPC")
	sum := int32(0)
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			average := float32(sum) / float32(count)
			return stream.SendAndClose(&calculatorPb.ComputeAverageResponse{
				Result: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		sum += req.GetNumber()
		count++
	}
}

func main() {
	log.Println("Calculator Server")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listent: %v", err)
	}
	s := grpc.NewServer()
	calculatorPb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v",err)
	}
}
