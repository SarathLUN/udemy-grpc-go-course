- generate TLS files:

```shell
./instructions.sh
```

- generate gRPC

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greetPb/greet.proto
```

- run server

```shell
go run greetServer/server.go
2021/09/26 01:02:19 /Users/sarath/go-workspace/src/github.com/SarathLUN/udemy-grpc-go-course <nil>
2021/09/26 01:02:19 starting GreetWithSsl server...
2021/09/26 01:03:05 GreetWithSsl function was invoked with greeting:{first_name:"Tony" last_name:"Stark"}
2021/09/26 01:03:08 GreetWithSsl function was invoked with greeting:{first_name:"Tony" last_name:"Stark"}

```

- run client

```shell
go run greetClient/client.go
2021/09/26 01:03:05 /Users/sarath/go-workspace/src/github.com/SarathLUN/udemy-grpc-go-course <nil>
2021/09/26 01:03:05 Hello, I'm a client
2021/09/26 01:03:05 Starting to do a UnaryWithSsl RPC...
2021/09/26 01:03:08 Response from GreetWithSsl: Hello, Tony
2021/09/26 01:03:08 Starting to do a UnaryWithSsl RPC...
2021/09/26 01:03:09 timeout was hit!

```