- `calculator.proto`

```protobuf
syntax = "proto3";

package calculatorPb;

option go_package = "github.com/SarathLUN/udemy-grpc-go-course/calculator-find-maximum/calculatorPb;calculatorPb";

message SquareRootRequest{
    int32 number = 1;
}

message SquareRootResponse{
    double number_root = 1;
}

service CalculatorService {
    rpc SquareRoot(SquareRootRequest) returns (SquareRootResponse){};
}
```

- generate protocol files

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculatorPb/calculator.proto
```

- start server

```shell
go run calculatorServer/server.go
2021/09/21 22:29:10 Calculator Server
2021/09/21 22:29:16 Received SquareRoot RPC: 10
2021/09/21 22:29:16 Received SquareRoot RPC: -2
2021/09/21 22:31:27 Received SquareRoot RPC: 34
2021/09/21 22:31:27 Received SquareRoot RPC: -4

```

- run client

```shell
go run calculatorClient/client.go
```

```shell
2021/09/21 22:29:16 Calculator Client
2021/09/21 22:29:16 Starting to do a SquareRoot Unary RPC...
2021/09/21 22:29:16 Result of square root of 10 = 3.1622776601683795
2021/09/21 22:29:16 Receive a negative number: -2
2021/09/21 22:29:16 InvalidArgument
2021/09/21 22:29:16 We probably sent a negative number!
```

```shell
go run calculatorClient/client.go
```

```shell
2021/09/21 22:31:27 Calculator Client
2021/09/21 22:31:27 Starting to do a SquareRoot Unary RPC...
2021/09/21 22:31:27 Result of square root of 34 = 5.830951894845301
2021/09/21 22:31:27 Receive a negative number: -4
2021/09/21 22:31:27 InvalidArgument
2021/09/21 22:31:27 We probably sent a negative number!
```
