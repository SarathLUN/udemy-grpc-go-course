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
