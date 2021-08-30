package main

import (
	"context"
	"github.com/SarathLUN/udemy-grpc-go-course/long-greet/longGreetPb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	log.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer cc.Close()
	c := longGreetPb.NewGreetServiceClient(cc)
	doClientSteam(c)
}

func doClientSteam(c longGreetPb.GreetServiceClient) {
	log.Println("starting to do a client stream RPC...")

	requests := []*longGreetPb.LongGreetRequest{
		{
			Greeting: &longGreetPb.Greeting{
				FirstName: "Stemphane",
			},
				}, {
			Greeting: &longGreetPb.Greeting{
				FirstName: "Jonh",
			},
				}, {
			Greeting: &longGreetPb.Greeting{
				FirstName: "Lucy",
			},
				}, {
			Greeting: &longGreetPb.Greeting{
				FirstName: "Mark",
			},
				}, {
			Greeting: &longGreetPb.Greeting{
				FirstName: "Piper",
			},
				},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	// we iterate over a slice and send each message individually
	for _, req := range requests {
		log.Printf("sending request: %v",req)
		stream.Send(req)
		time.Sleep(1000*time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while close and receive: %v", err)
	}
	log.Printf("LongGreet Result: %v\n", res)

}
