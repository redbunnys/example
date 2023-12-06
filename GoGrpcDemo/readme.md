# grpc

# 快速开始

- [grpcQuickStart](https://grpc.io/docs/languages/go/quickstart/)
- [ProtocolBuffersStart](https://protobuf.dev/getting-started/gotutorial/)

## 相关包下载
[protoc](https://github.com/protocolbuffers/protobuf/releases)
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
```
go get google.golang.org/grpc
go get google.golang.org/protobuf

```

## 生成
```
protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     protoc/greet.proto
```