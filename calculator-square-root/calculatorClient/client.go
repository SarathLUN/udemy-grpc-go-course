package main

import (
	"context"
	"github.com/SarathLUN/udemy-grpc-go-course/calculator-square-root/calculatorPb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	log.Println("Calculator Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		_ = cc.Close()
	}(cc)
	c := calculatorPb.NewCalculatorServiceClient(cc)
	doErrorUnary(c)
}

func doErrorUnary(c calculatorPb.CalculatorServiceClient) {
	log.Println("Starting to do a SquareRoot Unary RPC...")
	// correct call
	doErrorCall(c, 34)

	// error call
	doErrorCall(c, -4)
}

func doErrorCall(c calculatorPb.CalculatorServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorPb.SquareRootRequest{Number: n})
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			log.Println(resErr.Message())
			log.Println(resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Error calling SquareRoot: %v", err)
			return
		}
	}
	log.Printf("Result of square root of %v = %v", n, res.GetNumberRoot())
}
