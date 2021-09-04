- `calculator.proto`

```protobuf
syntax = "proto3";

package calculatorPb;

option go_package = "github.com/SarathLUN/udemy-grpc-go-course/calculator-find-maximum/calculatorPb;calculatorPb";

message FindMaximumRequest {
    int32 number=1;
}
message FindMaximumResponse {
    int32 maximum=1;
}

service CalculatorService {
    rpc FindMaximum(stream FindMaximumRequest) returns (stream FindMaximumResponse) {};
}
```

- generate protocol files

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculatorPb/calculator.proto
```

- start server

```shell
go run calculatorServer/server.go                           
2021/09/04 18:21:29 Calculator Server
2021/09/04 18:21:54 Received FindMaximum RPC

```

- run client

```shell
go run calculatorClient/client.go 
2021/09/04 18:21:54 Calculator Client
2021/09/04 18:21:54 start FindMaximum BiDi streaming RPC...
2021/09/04 18:21:54 Sending number: 4
2021/09/04 18:21:54 Received a new maximum: 4
2021/09/04 18:21:55 Sending number: 5
2021/09/04 18:21:55 Received a new maximum: 5
2021/09/04 18:21:56 Sending number: 7
2021/09/04 18:21:56 Received a new maximum: 7
2021/09/04 18:21:57 Sending number: 3
2021/09/04 18:21:58 Sending number: 7
2021/09/04 18:21:59 Sending number: 9
2021/09/04 18:21:59 Received a new maximum: 9
2021/09/04 18:22:00 Sending number: 23
2021/09/04 18:22:00 Received a new maximum: 23
2021/09/04 18:22:01 Sending number: 58
2021/09/04 18:22:01 Received a new maximum: 58

```
