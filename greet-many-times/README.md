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
    rpc DoGreetManyTimes(GreetManyTimesRequest) returns (steam GreetManyTimesResponse);
}

```

- in this case, I called it `DoGreetManyTimes`
- now generate protocol
```shell
>protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greet-many-times/greetManyTimesPB/greetManyTimes.proto
```
- now we implement server, then test running server
```shell
>go run greet-many-times/greetManyTimesServer/server.go 
2021/08/29 13:04:43 Starting GreetManyTimes Server ...
2021/08/29 13:04:43 listening on port: 50051

```
- now, implement client and run:
```shell
>go run greet-many-times/greetManyTimesClient/client.go
2021/08/29 20:20:19 Hello, I'm client. Let's do server streaming RPC.
2021/08/29 20:20:19 created connection client: &{%!f(*grpc.ClientConn=&{0xc000080980 0x1087540 localhost:50051 {passthrough  localhost:50051} localhost:50051 {<nil> <nil> [] [] <nil> <nil> {{1000000000 1.6 0.2 120000000000}} false false true 0 <nil>  {grpc-go/1.40.0 <nil> false [] <nil> <nil> {0 0 false} <nil> 0 0 32768 32768 0 <nil> true} [] <nil> 0 false true false <nil> <nil> <nil> <nil> []} 0xc0000744a0 {<nil> <nil> <nil> 0 grpc-go/1.40.0 {passthrough  localhost:50051}} 0xc0001ac420 {{{0 0} 0 0 0 0} 0xc000010260} {{0 0} 0 0 0 0} 0xc000032900 0xc00008a460 map[0xc0000d9080:{}] {0 0 false} pick_first 0xc00008a4b0 {<nil>} 0xc000074480 0 0xc000030280 {0 0} <nil>})}
2021/08/29 20:20:19 Result from GreetManyTimes: Hello, Tony, #0
2021/08/29 20:20:21 Result from GreetManyTimes: Hello, Tony, #1
2021/08/29 20:20:22 Result from GreetManyTimes: Hello, Tony, #2
2021/08/29 20:20:23 Result from GreetManyTimes: Hello, Tony, #3
2021/08/29 20:20:24 Result from GreetManyTimes: Hello, Tony, #4
2021/08/29 20:20:25 Result from GreetManyTimes: Hello, Tony, #5
2021/08/29 20:20:26 Result from GreetManyTimes: Hello, Tony, #6
2021/08/29 20:20:27 Result from GreetManyTimes: Hello, Tony, #7
2021/08/29 20:20:28 Result from GreetManyTimes: Hello, Tony, #8
2021/08/29 20:20:29 Result from GreetManyTimes: Hello, Tony, #9

```
- at the server side will get revoke logged:
```shell
2021/08/29 20:20:13 Starting GreetManyTimes Server ...
2021/08/29 20:20:13 listening on port: 50051
2021/08/29 20:20:19 DoGreetManyTimes function was invoked with greeting:{first_name:"Tony"  last_name:"Stark"}

```


