- generate gRPC

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greetPb/greet.proto
```

- run server

```shell
go run greetServer/server.go                                
2021/09/04 05:53:38 starting GreetEveryone server...
2021/09/04 05:54:27 GreetEveryone is invoked

```

- run client

```shell
go run greetClient/client.go 
2021/09/04 05:54:27 Hello, I'm a client
2021/09/04 05:54:27 Starting to do a BiDi streaming RPC...
2021/09/04 05:54:27 Sending message: greeting:{first_name:"Stephane"}
2021/09/04 05:54:27 Received: Hello, Stephane! 
2021/09/04 05:54:28 Sending message: greeting:{first_name:"John"}
2021/09/04 05:54:28 Received: Hello, John! 
2021/09/04 05:54:29 Sending message: greeting:{first_name:"Lucy"}
2021/09/04 05:54:29 Received: Hello, Lucy! 
2021/09/04 05:54:30 Sending message: greeting:{first_name:"Mark"}
2021/09/04 05:54:30 Received: Hello, Mark! 
2021/09/04 05:54:31 Sending message: greeting:{first_name:"Piper"}
2021/09/04 05:54:31 Received: Hello, Piper! 

```