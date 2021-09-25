package main

import (
	"context"
	"github.com/SarathLUN/udemy-grpc-go-course/GreetWithSsl/greetPb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"time"
)

type server struct {
	greetPb.UnimplementedGreetServiceServer
}

func (s server) GreetWithSsl(ctx context.Context, req *greetPb.GreetWithSslRequest) (*greetPb.GreetWithSslResponse, error) {
	log.Printf("GreetWithSsl function was invoked with %v", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// the client canceled the request
			log.Println("the client cancel request.")
			return nil, status.Error(codes.Canceled, "the client canceled the request.")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello, " + firstName
	res := &greetPb.GreetWithSslResponse{
		Result: result,
	}
	return res, nil
}
func main() {
	log.Println(os.Getwd())
	log.Println("starting GreetWithSsl server...")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// working on ssl
	tls := true
	var opts []grpc.ServerOption
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("fail to load certificates: %v", err)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)
	greetPb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
