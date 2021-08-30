package main

import (
	"context"
	calculator "github.com/SarathLUN/udemy-grpc-go-course/calculator-compute-average/calculatorpb"
	"google.golang.org/grpc"
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
	doClientStreaming(c)
}

func doClientStreaming(c calculator.CalculatorServiceClient) {
	log.Println("starting to do a ComputeAverage client streaming RPC...")
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}
	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		log.Printf("sending number: %v", number)
		err := stream.Send(&calculator.ComputeAverageRequest{Number: number})
		if err != nil {
			log.Fatalf("failed to send the client stream: %v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to close and receive: %v", err)
	}
	log.Printf("The average: %v", res.GetResult())
}
