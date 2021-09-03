package main

import (
	"context"
	"github.com/SarathLUN/udemy-grpc-go-course/GreetEveryone/greetPb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	log.Println("Hello, I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := greetPb.NewGreetServiceClient(cc)
	doBiDiStreaming(c)
}

func doBiDiStreaming(c greetPb.GreetServiceClient) {
	log.Println("Starting to do a BiDi streaming RPC...")
	// we create the stream by invoking client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}
	requests := []*greetPb.GreetEveryoneRequest{
		{
			Greeting: &greetPb.Greeting{
				FirstName: "Stephane",
			},
		}, {
			Greeting: &greetPb.Greeting{
				FirstName: "John",
			},
		}, {
			Greeting: &greetPb.Greeting{
				FirstName: "Lucy",
			},
		}, {
			Greeting: &greetPb.Greeting{
				FirstName: "Mark",
			},
		}, {
			Greeting: &greetPb.Greeting{
				FirstName: "Piper",
			},
		},
	}
	waitc := make(chan struct{})
	// we send a bunch of messages to the server (go routine)
	go func() {
		for _, req := range requests {
			log.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the server (go routine)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving: %v", err)
				break
			}
			log.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()
	// block until everything done
	<-waitc
}
