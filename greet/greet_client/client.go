package main

import (
	"github.com/SarathLUN/udemy-grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.Println("Hello, I'm client.")
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
	client := greetpb.NewGreetServiceClient(conn)
	log.Printf("created connection client: %f",client)
}
