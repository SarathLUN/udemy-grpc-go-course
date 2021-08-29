# Server Stream

- define `proto` definition
- copy from the previous `greet.proto` to `greetManyTimes.proto` and add one more `rpc` function

```protobuf
[...]

message GreetManyTimesRequest{
    Greeting greeting = 1;
}

message GreetManyTimesResponse{
    string result = 1;
}

service GreetService{
    rpc DoGreet(GreetRequest) returns (GreetResponse);
    
        // server streaming
    rpc DoGreetManyTimes(GreetManyTimesRequest) returns (GreetManyTimesResponse);
}

```

- in this case, I called it `DoGreetManyTimes`
- now generate protocol
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greet-many-times/greetManyTimesPB/greetManyTimes.proto
```