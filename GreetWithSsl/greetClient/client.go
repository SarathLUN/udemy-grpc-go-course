package main

import (
	"context"
	"github.com/SarathLUN/udemy-grpc-go-course/GreetWithSsl/greetPb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"time"
)

func main() {
	log.Println(os.Getwd())
	log.Println("Hello, I'm a client")
	tls := true
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("fail to load certificate: %v", err)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := greetPb.NewGreetServiceClient(cc)

	// this call should complete
	doUnaryWithSsl(c, 5*time.Second)

	// this call should be timeout
	doUnaryWithSsl(c, 1*time.Second)
}

func doUnaryWithSsl(c greetPb.GreetServiceClient, seconds time.Duration) {
	log.Println("Starting to do a UnaryWithSsl RPC...")
	req := &greetPb.GreetWithSslRequest{
		Greeting: &greetPb.Greeting{
			FirstName: "Tony",
			LastName:  "Stark",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), seconds)
	defer cancel()

	res, err := c.GreetWithSsl(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			// this error response from grpc
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("timeout was hit!")
			} else {
				log.Printf("unexpected error: %v", statusErr)
			}
		} else {
			// this error response from go
			log.Fatalf("error while calling GreetWithSsl RPC: %v", err)
		}
		return
	}
	log.Printf("Response from GreetWithSsl: %v", res.Result)
}
