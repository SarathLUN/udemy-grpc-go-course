package main

import (
	"context"
	"github.com/SarathLUN/udemy-grpc-go-course/GreetWithDeadline/greetPb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	// this call should complete
	doUnaryWithDeadline(c, 5*time.Second)

	// this call should be timeout
	doUnaryWithDeadline(c, 1*time.Second)
}

func doUnaryWithDeadline(c greetPb.GreetServiceClient, seconds time.Duration) {
	log.Println("Starting to do a UnaryWithDeadline RPC...")
	req := &greetPb.GreetWithDeadlineRequest{
		Greeting: &greetPb.Greeting{
			FirstName: "Tony",
			LastName:  "Stark",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), seconds)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			// this error response from grpc
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("timeout was hit! deadline was exceeded")
			} else {
				log.Printf("unexpected error: %v", statusErr)
			}
		} else {
			// this error response from go
			log.Fatalf("error while calling GreetWithDeadline RPC: %v", err)
		}
		return
	}
	log.Printf("Response from GreetWithDeadline: %v", res.Result)
}
