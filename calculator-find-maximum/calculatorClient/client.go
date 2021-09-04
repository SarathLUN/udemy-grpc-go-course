package main

import (
	"context"
	"github.com/SarathLUN/udemy-grpc-go-course/calculator-find-maximum/calculatorPb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	log.Println("Calculator Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := calculatorPb.NewCalculatorServiceClient(cc)
	doBiDiStreaming(c)
}

func doBiDiStreaming(c calculatorPb.CalculatorServiceClient) {
	log.Println("start FindMaximum BiDi streaming RPC...")
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("error while opening stream and calling FindMaximum: %v", err)
	}
	waitc := make(chan struct{})
	// send go routine
	go func() {
		numbers := []int32{4, 5, 7, 3, 7, 9, 23, 58}
		for _, number := range numbers {
			log.Printf("Sending number: %v", number)
			stream.Send(&calculatorPb.FindMaximumRequest{Number: number})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("problem while reading server stream: %v", err)
			}
			maximum := res.GetMaximum()
			log.Printf("Received a new maximum: %v\n", maximum)
		}
		close(waitc)
	}()
	// close channel
	<-waitc

}
