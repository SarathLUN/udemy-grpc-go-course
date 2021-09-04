package main

import (
	"github.com/SarathLUN/udemy-grpc-go-course/calculator-find-maximum/calculatorPb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
	calculatorPb.UnimplementedCalculatorServiceServer
}

func (s server) FindMaximum(stream calculatorPb.CalculatorService_FindMaximumServer) error  {
	// log function invoked
	log.Println("Received FindMaximum RPC")
	// declare max
	var maximum int32
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
			return err
		}
		// get value from client request
		number := req.GetNumber()
		if number > maximum {
			maximum = number
			err = stream.Send(&calculatorPb.FindMaximumResponse{Maximum: maximum})
			if err != nil {
				log.Fatalf("error while sending data to client: %v",err)
				return err
			}
		}

	}
}

func main() {
	// console message server start
	log.Println("Calculator Server")
	// create listener
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v",err)
	}
	// create server
	s := grpc.NewServer()
	calculatorPb.RegisterCalculatorServiceServer(s, &server{})
	// serve
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v",err)
	}
}
