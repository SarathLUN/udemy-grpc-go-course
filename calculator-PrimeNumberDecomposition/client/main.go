package main

import (
	"context"
	calculator "github.com/SarathLUN/udemy-grpc-go-course/calculator-PrimeNumberDecomposition/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {

	log.Println("Calculator Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	c := calculator.NewCalculatorServiceClient(cc)
	doServerStream(c)
}

func doServerStream(c calculator.CalculatorServiceClient) {
	log.Println("starting to do a PrimeNumberDecomposition, server stream RPC...")
	req := &calculator.PrimeNumberDecompositionRequest{Number: 210}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}
		log.Printf("result server stream: %v", res.GetPrimeFactor())
	}
}
