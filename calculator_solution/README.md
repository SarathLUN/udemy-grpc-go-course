- start server
```shell
gor calculator_server/server.go
2021/08/27 07:13:17 Calculator Server, port:50051
2021/08/27 07:13:17 Successful: listener started!
2021/08/27 07:14:00 Server got request(*calculatorpb.SumRequest): first_number:6 second_number:4
2021/08/27 07:14:17 Server got request(*calculatorpb.SumRequest): first_number:6 second_number:4
2021/08/27 07:15:48 Server got request(*calculatorpb.SumRequest): first_number:6 second_number:4

```
- run client
```shell
gor calculator_solution/calculator_client/client.go 
2021/08/27 07:14:00 Calculator Client
2021/08/27 07:14:00 start do Sum (Unary RPC)
2021/08/27 07:14:00 Result: result:10

```

```shell
gor calculator_solution/calculator_client/client.go
2021/08/27 07:14:17 Calculator Client
2021/08/27 07:14:17 start do Sum (Unary RPC)
2021/08/27 07:14:17 Result: 10

```

```shell
gor calculator_solution/calculator_client/client.go
2021/08/27 07:15:48 Calculator Client
2021/08/27 07:15:48 start do Sum (Unary RPC)
2021/08/27 07:15:48 Result: 10

```