package main

import (
	"context"
	calculate "github.com/SarathLUN/udemy-grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.Println("Hello, I'm client.")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := calculate.NewSumServiceClient(conn)
	req := &calculate.SumRequest{Input: &calculate.Input{Num1: 3, Num2: 4}}
	res, err := client.DoSum(context.Background(), req)
	if err != nil {
		log.Fatalf("fail DoSum: %v", err)
	}
	log.Printf("response: %v", res.Result)
}
