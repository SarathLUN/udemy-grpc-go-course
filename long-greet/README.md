# this will demonstrate Client Streaming RPC

## define `proto` file

As usual, we start from define `proto` file and generate gRPC files

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative long-greet/longGreetPb/longGreet.proto
```

- implement server to listen the client stream

```go
[...]

func (s server) LongGreet(stream longGreetPb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with a streaming request")
	var result string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finish reading the client stream
			return stream.SendAndClose(&longGreetPb.LongGreetResponse{Result: result})
		}
		if err != nil {
			log.Fatalf("failed reading client stream: %v\n", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "!\n"
	}
	return nil
}

[...]

```
## now start the server
```shell
go run long-greet/longGreetServer/main.go
2021/08/30 21:51:33 starting LongGreet Server
2021/08/30 21:51:40 LongGreet function was invoked with a streaming request
```
## now run the client
```shell
go run long-greet/longGreetClient/main.go
2021/08/30 21:51:40 Hello I'm a client
2021/08/30 21:51:40 starting to do a client stream RPC...
2021/08/30 21:51:40 sending request: greeting:{first_name:"Stemphane"}
2021/08/30 21:51:41 sending request: greeting:{first_name:"Jonh"}
2021/08/30 21:51:42 sending request: greeting:{first_name:"Lucy"}
2021/08/30 21:51:43 sending request: greeting:{first_name:"Mark"}
2021/08/30 21:51:44 sending request: greeting:{first_name:"Piper"}
2021/08/30 21:51:45 LongGreet Result: result:"Hello Stemphane!, Hello Jonh!, Hello Lucy!, Hello Mark!, Hello Piper!, "
```