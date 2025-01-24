## Meta 仓库地址

用来存放 `proto` 文件仓库，对应着仓库 `MetaCode` 。


下载 protoc-gen-go and protoc-gen-go-grpc:
```shell
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

使用 

```shell
protoc -I . --go_out=./pb/helloWorldPB --go-grpc_out=./pb/helloWorldPB  .\pb\helloWorldPB\hello_world.proto
```