```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculator-PrimeNumberDecomposition/calculatorpb/calculator.proto
```

- start server

```shell
go run calculator-PrimeNumberDecomposition/server/main.go 
2021/08/29 23:07:36 Calculator Server
2021/08/29 23:07:57 Received PrimeNumberDecomposition RPC: number:210
2021/08/29 23:07:57 Divisor has increased to 3
2021/08/29 23:07:57 Divisor has increased to 4
2021/08/29 23:07:57 Divisor has increased to 5
2021/08/29 23:07:57 Divisor has increased to 6
2021/08/29 23:07:57 Divisor has increased to 7

```

- run test client PrimeNumberDecomposition of `210`:

```shell
go run calculator-PrimeNumberDecomposition/client/main.go 
2021/08/29 23:07:57 Calculator Client
2021/08/29 23:07:57 starting to do a PrimeNumberDecomposition, server stream RPC...
2021/08/29 23:07:57 result server stream: 2
2021/08/29 23:07:57 result server stream: 3
2021/08/29 23:07:57 result server stream: 5
2021/08/29 23:07:57 result server stream: 7

```
