In this exercise, your goal is to implement a **_ComputeAverage RPC Client Streaming API in a CalculatorService_**:

- The function takes a stream of Request message that has one integer, and returns a Response with a double that
  represents the computed average
- Remember to first implement the service definition in a .proto file, alongside the RPC messages
- Implement the Server code first
- Test the server code by implementing the Client

_Example:_

The client will send a stream of number (1,2,3,4) and the server will respond with (2.5), because (1+2+3+4)/4 = 2.5

Good luck!

---

# My Solution:

- define `proto` file

```protobuf
syntax = "proto3";

package calculatorPb;

option go_package = "github.com/SarathLUN/udemy-grpc-go-course/calculator-compute-average/calculatorPb;calculatorPb";

message ComputeAverageRequest {
    int32 number = 1;
}

message ComputeAverageResponse {
    float result = 1;
}

service CalculatorService {
    rpc ComputeAverage(stream ComputeAverageRequest) returns (ComputeAverageResponse) {};
}
```

- generate gRPC

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculator-compute-average/calculatorPb/calculator.proto
```

- implement server

```go
[...]

func (s server) ComputeAverage(stream calculatorPb.CalculatorService_ComputeAverageServer) error{
	log.Println("Received ComputeAverage RPC")
	sum := int32(0)
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			average := float32(sum) / float32(count)
			return stream.SendAndClose(&calculatorPb.ComputeAverageResponse{
				Result: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		sum += req.GetNumber()
		count++
	}
}

[...]
```

- implement client

```go

[...]

func doClientStreaming(c calculator.CalculatorServiceClient) {
	log.Println("starting to do a ComputeAverage client streaming RPC...")
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}
	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		log.Printf("sending number: %v", number)
		err := stream.Send(&calculator.ComputeAverageRequest{Number: number})
		if err != nil {
			log.Fatalf("failed to send the client stream: %v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to close and receive: %v", err)
	}
	log.Printf("The average: %v", res.GetResult())
}

```
- start the server
```shell
go run calculator-compute-average/calculatorServer/server.go 
2021/08/30 23:44:36 Calculator Server
2021/08/30 23:44:55 Received ComputeAverage RPC

```
- run the client
```shell
go run calculator-compute-average/calculatorClient/client.go 
2021/08/30 23:44:55 Calculator Client
2021/08/30 23:44:55 starting to do a ComputeAverage client streaming RPC...
2021/08/30 23:44:55 sending number: 3
2021/08/30 23:44:55 sending number: 5
2021/08/30 23:44:55 sending number: 9
2021/08/30 23:44:55 sending number: 54
2021/08/30 23:44:55 sending number: 23
2021/08/30 23:44:55 The average: 18.8

```
