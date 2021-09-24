- generate gRPC

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greetPb/greet.proto
```

- run server

```shell
go run greetServer/server.go
2021/09/24 19:12:02 starting GreetWithDeadline server...
2021/09/24 19:12:13 GreetWithDeadline function was invoked with greeting:{first_name:"Tony"  last_name:"Stark"}
2021/09/24 19:12:16 GreetWithDeadline function was invoked with greeting:{first_name:"Tony"  last_name:"Stark"}

```

- run client

```shell
go run greetClient/client.go
2021/09/24 19:12:13 Hello, I'm a client
2021/09/24 19:12:13 Starting to do a UnaryWithDeadline RPC...
2021/09/24 19:12:16 Response from GreetWithDeadline: Hello, Tony
2021/09/24 19:12:16 Starting to do a UnaryWithDeadline RPC...
2021/09/24 19:12:17 timeout was hit! deadline was exceeded

```