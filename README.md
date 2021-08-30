# Getting started
```shell
go mod init
cat go.mod
```
output:
```shell
module github.com/SarathLUN/udemy-grpc-go-course

go 1.17
```
then install grpc
```shell
go get -v google.golang.org/grpc
```

output:
```shell
google.golang.org/protobuf/internal/set
google.golang.org/protobuf/internal/flags
google.golang.org/protobuf/internal/pragma
google.golang.org/protobuf/internal/detrand
google.golang.org/protobuf/internal/version
google.golang.org/protobuf/internal/errors
google.golang.org/protobuf/encoding/protowire
google.golang.org/protobuf/reflect/protoreflect
google.golang.org/protobuf/internal/mapsort
google.golang.org/protobuf/reflect/protoregistry
google.golang.org/protobuf/internal/fieldsort
google.golang.org/protobuf/internal/descopts
google.golang.org/protobuf/internal/descfmt
google.golang.org/protobuf/runtime/protoiface
google.golang.org/protobuf/internal/strs
google.golang.org/protobuf/internal/genid
google.golang.org/protobuf/internal/encoding/text
google.golang.org/protobuf/internal/encoding/messageset
google.golang.org/protobuf/proto
google.golang.org/protobuf/internal/encoding/defval
google.golang.org/protobuf/encoding/prototext
google.golang.org/protobuf/internal/filedesc
google.golang.org/protobuf/internal/encoding/tag
google.golang.org/protobuf/internal/impl
google.golang.org/protobuf/internal/filetype
google.golang.org/protobuf/runtime/protoimpl
google.golang.org/protobuf/types/known/durationpb
google.golang.org/protobuf/types/known/anypb
google.golang.org/protobuf/types/known/timestamppb
github.com/golang/protobuf/proto
github.com/golang/protobuf/ptypes/duration
github.com/golang/protobuf/ptypes/any
github.com/golang/protobuf/ptypes/timestamp
google.golang.org/grpc/encoding/proto
google.golang.org/grpc/credentials
google.golang.org/grpc/binarylog/grpc_binarylog_v1
github.com/golang/protobuf/ptypes
google.golang.org/genproto/googleapis/rpc/status
google.golang.org/grpc/resolver
google.golang.org/grpc/peer
google.golang.org/grpc/internal/channelz
google.golang.org/grpc/internal/status
google.golang.org/grpc/internal/metadata
google.golang.org/grpc/balancer/grpclb/state
google.golang.org/grpc/internal/transport/networktype
google.golang.org/grpc/internal
google.golang.org/grpc/internal/resolver/passthrough
google.golang.org/grpc/internal/grpcutil
google.golang.org/grpc/internal/resolver/unix
google.golang.org/grpc/internal/resolver/dns
google.golang.org/grpc/status
google.golang.org/grpc/balancer
google.golang.org/grpc/internal/binarylog
google.golang.org/grpc/internal/serviceconfig
google.golang.org/grpc/balancer/base
google.golang.org/grpc/internal/transport
google.golang.org/grpc/internal/resolver
google.golang.org/grpc/balancer/roundrobin
google.golang.org/grpc
go get: added github.com/golang/protobuf v1.4.3
go get: added golang.org/x/net v0.0.0-20200822124328-c89045814202
go get: added golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd
go get: added golang.org/x/text v0.3.0
go get: added google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
go get: added google.golang.org/grpc v1.40.0
go get: added google.golang.org/protobuf v1.25.0
```
check `go.mod` again
```shell
cat go.mod
```
output:
```shell
module github.com/SarathLUN/udemy-grpc-go-course

go 1.17

require (
	github.com/golang/protobuf v1.4.3 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.40.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)
```
# start the proto file
- when try to use `option go_package = ".;greetpb";` GoLand IDE got error `Built-in option 'go_package' not found`
- this error is gone once change to use full path
```protobuf
option go_package = "github.com/SarathLUN/udemy-grpc-go-course;greetpb";
```
- before generate gRPC
```shell
tree .
.
├── README.md
├── calculator
├── go.mod
├── go.sum
├── greet
│   └── greetpb
└── greet.proto

3 directories, 4 files
```
- generate gPRC
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greet/greetpb/greet.proto
```
- after generated gPRC
```shell
tree .

.
├── README.md
├── calculator
├── go.mod
├── go.sum
└── greet
    └── greetpb
        ├── greet.pb.go
        ├── greet.proto
        └── greet_grpc.pb.go

3 directories, 6 files
```
# start the server
- when follow the video, got error of strut `server`
```shell
./server.go:22:40: cannot use &server{} (type *server) as type greetpb.GreetServiceServer in argument to greetpb.RegisterGreetServiceServer:
	*server does not implement greetpb.GreetServiceServer (missing greetpb.mustEmbedUnimplementedGreetServiceServer method)
```
- we need to implement the interface, so struct `server` need to have code like this:
```go
type server struct {
	greetpb.UnimplementedGreetServiceServer
}
```
- try to run server again, successful:
```shell
go run greet/greet_server/server.go
2021/08/24 00:20:49 Hello world!

```
# start the client
- make sure the server is running
- then we can run test the client, and it worked!
```shell
gor greet_client/client.go
2021/08/24 23:25:09 Hello, I'm client.
2021/08/24 23:25:09 created connection client: &{%!f(*grpc.ClientConn=&{0xc00014e940 0x11ba340 localhost:50051 {passthrough  localhost:50051} localhost:50051 {<nil> <nil> [] [] <nil> <nil> {{1000000000 1.6 0.2 120000000000}} false false true 0 <nil>  {grpc-go/1.40.0 <nil> false [] <nil> <nil> {0 0 false} <nil> 0 0 32768 32768 0 <nil> true} [] <nil> 0 false true false <nil> <nil> <nil> <nil> []} 0xc000156480 {<nil> <nil> <nil> 0 grpc-go/1.40.0 {passthrough  localhost:50051}} 0xc0001ce420 {{{0 0} 0 0 0 0} 0xc000136250} {{0 0} 0 0 0 0} 0xc000138840 0xc000144410 map[0xc000175080:{}] {0 0 false} pick_first 0xc000144460 {<nil>} 0xc000156460 0 0xc000142240 {0 0} <nil>})}

```

---

# gRPC Unary API
- our message is `Greeting` and contains `first_name` and `last_name` string field
- it will take a `GreetRequest` that contains a `Greeting`
- it will return a `GreetResponse` that contains a result string
- after update `proto` file, we need to re-generate again
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greet/greetpb/greet.proto

ll greet/greetpb

total 40
-rw-r--r--  1 sarath  staff   9.2K Aug 25 00:13 greet.pb.go
-rw-r--r--  1 sarath  staff   359B Aug 25 00:08 greet.proto
-rw-r--r--  1 sarath  staff   3.5K Aug 25 00:13 greet_grpc.pb.go
```
- new files generated successfully
- now implement server and run test
```go
func (s server) DoGreet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := in.GetGreeting().GetFirstName()
	result := "Hello, " + firstName
	res := &greetpb.GreetResponse{Result: result}
	return res, nil
}

```
```shell
go run greet/greet_server/server.go
2021/08/25 00:24:09 Hello world!

```
- now implement and run the client
```shell
go run greet/greet_client/client.go
2021/08/25 00:36:32 Hello, I'm client.
2021/08/25 00:36:32 created connection client: &{%!f(*grpc.ClientConn=&{0xc00014e980 0x1087540 localhost:50051 {passthrough  localhost:50051} localhost:50051 {<nil> <nil> [] [] <nil> <nil> {{1000000000 1.6 0.2 120000000000}} false false true 0 <nil>  {grpc-go/1.40.0 <nil> false [] <nil> <nil> {0 0 false} <nil> 0 0 32768 32768 0 <nil> true} [] <nil> 0 false true false <nil> <nil> <nil> <nil> []} 0xc000156480 {<nil> <nil> <nil> 0 grpc-go/1.40.0 {passthrough  localhost:50051}} 0xc0001ce420 {{{0 0} 0 0 0 0} 0xc000138250} {{0 0} 0 0 0 0} 0xc00013a840 0xc000144410 map[0xc000175080:{}] {0 0 false} pick_first 0xc000144460 {<nil>} 0xc000156460 0 0xc000142240 {0 0} <nil>})}
2021/08/25 00:36:33 response form greet: Hello, Tony
```
- at the serve side we also see the log request:
```shell
2021/08/25 00:36:33 Greet function was invoked with greeting:{first_name:"Tony" last_name:"Stark"}
```
# Exercises:
- my own code [here](/calculator).
- follow solution [here](/calculator_solution).

# Server Streaming RPC

- class work [here](/greet-many-times)
- exercise [here](/calculator-PrimeNumberDecomposition)

