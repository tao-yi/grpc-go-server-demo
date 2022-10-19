## Install Protoc

- [Guick Start | Go](https://grpc.io/docs/languages/go/quickstart/)

```shell
$ go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
$ go get google.golang.org/protobuf/cmd/protoc-gen-go
$ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Update your PATH so that the protoc compiler can find the plugins:
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Package
https://developers.google.com/protocol-buffers/docs/proto3#packages
You can add a optional package specifier to a `.proto` file to prevent name clashes between protocol message types
```
package foo.bar;

message Open {
}
```
You can then use the packages specifier when defining fields of your message type:
```
message Foo {
    foo.bar.Open open = 1;
}
```
The way a package specifier affects the generated code depends on you chosen language:
- In Go, the package is used as the Go package name, unless you explicitly provide an `option go_package` in your `.proto` file
- In Python, the package directive is ignored, since Python modules are organized according to their location in the file system.

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

grpcurl -plaintext localhost:50051 user.service.v1.UserService/GetUserInfo
grpcurl -plaintext localhost:50051 user.UserService/GetArticles
```

## import googleapis

https://github.com/googleapis/googleapis/tree/master/google/api

## Install [Evans](https://github.com/ktr0731/evans#macos)

```shell
$ brew tap ktr0731/evans
$ brew install evans
```

If your server is enabling gRPC reflection, you can launch Evans with only -r (--reflection) option.

```shell
$ evans -r repl
```
