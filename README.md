## Install Protoc

- [Guick Start | Go](https://grpc.io/docs/languages/go/quickstart/)

```shell
$ go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
$ go get google.golang.org/protobuf/cmd/protoc-gen-go
$ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Update your PATH so that the protoc compiler can find the plugins:
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Generate gRPC code

```shell
$ protoc \
    --proto_path=./proto \
    --go_out=. \
    --go-grpc_out=. \
    ./**/*.proto
```

- `--proto_path` 或者 `-I` 参数用以指定 proto 文件的位置
- `--go_out` 用来指定 go 代码目录结构的生成位置

  - `paths` 有两个选项：`source_relative` 和 `import`，默认值是 `import`。
  - `import` 表示用 `option go_package` 参数指定的包结构来创建目录层级
  - `source_relative` 表示在 proto 源文件的当前目录下创建目录层级

- `option go_package = "./;pb";` 表示生成代码到 out 目录下的 `./` 目录中，并且使用 `package pb`
- 搭配 `--go_out=./pb` 说明在 `./pb` 文件夹下的 `./` 目录生成代码，并使用 `package pb`

`protoc` 会生成两套代码：

- Code for populating, serializing, and retrieving HelloRequest and HelloReply message types.
- Generated client and server code.

## Install [grpcurl](https://github.com/fullstorydev/grpcurl)

```shell
brew install grpcurl

# list registered services
grpcurl -plaintext localhost:50051 list

grpcurl -plaintext localhost:50051 v1.UserService/GetUserInfo
grpcurl -plaintext localhost:50051 user.UserService/GetArticles
```

## import googleapis

https://github.com/googleapis/googleapis/tree/master/google/api
