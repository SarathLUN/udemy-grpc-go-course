```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculator/calculatorpb/calculate.proto
```

- start server

```shell
>go run calculator/server/server.go
2021/08/25 23:06:04 starting server on port 50051...
2021/08/25 23:06:04 server is running on 50051

```

- run test client:

```shell
>go run calculator/client/client.go 
2021/08/25 23:06:19 Hello, I'm client.
2021/08/25 23:06:19 response: 7

```
