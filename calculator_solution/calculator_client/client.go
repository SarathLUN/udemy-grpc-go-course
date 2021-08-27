package main

import (
	"context"
	calculate "github.com/SarathLUN/udemy-grpc-go-course/calculator_solution/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {

	log.Println("Calculator Client")
	// create client connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: dial: %v", err)
	}
	defer cc.Close()
	// create client
	c := calculate.NewDoSumClient(cc)
	doUnary(c)
}

func doUnary(c calculate.DoSumClient) {
	log.Println("start do Sum (Unary RPC)")
	req := &calculate.SumRequest{SecondNumber: 3, FirstNumber: 6}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: Sum: %v",err)
	}
	log.Printf("Result: %v",res.Result)
}
