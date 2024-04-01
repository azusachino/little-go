# Grpc

## define protobuf

```protobuf
syntax = "proto3";

option go_package = "github.com/azusachino/golong/practices/grpc/helloworld";
option java_multiple_files = true;
option java_package = "cn.az.code.grpc";
option java_outer_classname = "HelloWorldProto";

package helloworld;

service Greeter {
  rpc SayHello(HelloRequest) returns (HelloResponse) {}
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}
```

## install compiler

```shell
go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## compile

```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```
