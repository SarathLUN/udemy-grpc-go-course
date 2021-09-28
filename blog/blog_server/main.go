package main

import (
	blogpb "github.com/SarathLUN/udemy-grpc-go-course/blog/blog_pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

type server struct {
	blogpb.UnimplementedBlogServiceServer
}

func main() {
	// if we crush the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// starting messages
	log.Println("Blog Server started...")
	log.Println("Ctrl+C to stop the server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// wait for Ctrl+C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until a signal is received
	<-ch
	// stop the server
	log.Println("Stopping the server")
	s.Stop()
	// close the listener
	log.Println("Closing the listener")
	_ = lis.Close()
	// End of program
	log.Println("End of Program")

}
