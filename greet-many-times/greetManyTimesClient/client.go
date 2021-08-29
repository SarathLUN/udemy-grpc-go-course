package main

import (
	"context"
	greetManyTimesPb "github.com/SarathLUN/udemy-grpc-go-course/greet-many-times/greetManyTimesPB"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	log.Println("Hello, I'm client. Let's do server streaming RPC.")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial to server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("can't close the client connection: %v", err)
		}
	}(conn)
	client := greetManyTimesPb.NewGreetServiceClient(conn)
	log.Printf("created connection client: %f", client)
	req := &greetManyTimesPb.GreetManyTimesRequest{
		Greeting: &greetManyTimesPb.Greeting{
			FirstName: "Tony",
			LastName: "Stark",
		},
	}
	streamResult, err := client.DoGreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: call server streaming: %v", err)
	}
	for {
		msg, err := streamResult.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error: receiving message from server: %v", err)
		}
		log.Printf("Result from GreetManyTimes: %v",msg.GetResult())
	}
}


