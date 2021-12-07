# Todo-vue-grpc

![image-20211207141441940](/Users/kaiwen/Library/Application Support/typora-user-images/image-20211207141441940.png)

## 定义proto文件

## 通过proto文件生成go文件以及grpc调用相关文件（如果使用go-micro框架，还需要生成micro相关文件）

> [首先需要下载proto](--go_out=./proto/gen/go/todo/v1 --go_opt=paths=source_relative proto/idl/todo/v1/todo_api.proto proto/idl/todo/v1/todo.proto)
>
> 之后下载proto生成go和grpc的插件：
>
> ```
> $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
> $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
> ```

### 生成go文件：

```
protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto
```

### 生成grpc调用相关文件

```
// 同时生成go和grpc
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto
```

> ### **注(注意import引入)：**
>
> 如果proto文件存在相互引用的关系，则import中的路径要和执行生成语句时文件路径一致：
>
> #### eg:
>
> ##### 项目结构：
>
> ```
>  └── proto
>       └── idl
>           └── todo
>               └── v1
>                   ├── todo.proto
>                   └── todo_api.proto
> ```
>
> ##### todo_api.proto中import（引入）todo.proto：
>
> ```
> syntax="proto3";
> 
> package todo.v1;
> 
> option go_package = "github.com/LHaosss/todo-vue-grpc/proto/gen/go/todo/v1;todov1";
> // this
> import "proto/idl/todo/v1/todo.proto"; // this import content should be samilar to file which generating Go file depends on.
> ```
>
> 根目录执行生成命令：
>
> ```
> protoc --go_out=./proto/gen/todo/todo/v1 --go_opt=paths=source_relative proto/idl/todo/v1/todo_api.proto proto/idl/todo/v1/todo.proto
> ```

### 生成micro相关文件

#### 下载所需插件：

```
go get github.com/micro/protoc-gen-micro/v2
```

#### 生成命令

```
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
```

### 命令：

```
#!/usr/bin/env bash
set -eo pipefail

CUR=$(dirname $0)

GEN_OUT_DIR=${CUR}/gen

GEN_TOOL=${CUR}/protoc_gen_plugin.bash
PROTO_PATH=${CUR}/idl
PROTO_INCLUDE_PATH_1="${CUR}/idl"

# go
${GEN_TOOL} \
  --proto_path=${PROTO_PATH} \
  --proto_include_path=${PROTO_INCLUDE_PATH_1} \
  --plugin_name=go \
  --plugin_out=${LANG_OUT} \
  --plugin_opt=paths=source_relative
# grpc
${GEN_TOOL} \
  --proto_path=${PROTO_PATH} \
  --proto_include_path=${PROTO_INCLUDE_PATH_1} \
  --plugin_name=go-grpc \
  --plugin_out=${LANG_OUT} \
  --plugin_opt=paths=source_relative
# go-micro
${GEN_TOOL} \
  --proto_path=${PROTO_PATH} \
  --proto_include_path=${PROTO_INCLUDE_PATH_1} \
  --plugin_name=micro \
  --plugin_out=${LANG_OUT} \
  --plugin_opt=paths=source_relative
```

## go、grpc、micro文件：

#### go:

在go文件中主要包含结构的定义

#### Grpc:

grpc中则是定义了客户端接口和对应结构体，以及生成客户端对象的方法

服务端只有接口

#### micro：

其中则包含了客户端和服务端的接口以及结构体和对象生成，同时对服务端的方法进行了实现

### use gRPC with TLS in go

#### Go

##### Base case - no encryption or authentication

Client:

```go
conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
// error handling omitted
client := pb.NewGreeterClient(conn)
// ...
```

Server:

```go
s := grpc.NewServer()
lis, _ := net.Listen("tcp", "localhost:50051")
// error handling omitted
s.Serve(lis)
```

##### With server authentication SSL/TLS

Client:

```go
creds, _ := credentials.NewClientTLSFromFile(certFile, "")
conn, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
// error handling omitted
client := pb.NewGreeterClient(conn)
// ...
```

Server:

```go
creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)
s := grpc.NewServer(grpc.Creds(creds))
lis, _ := net.Listen("tcp", "localhost:50051")
// error handling omitted
s.Serve(lis)
```

##### Authenticate with Google

```go
pool, _ := x509.SystemCertPool()
// error handling omitted
creds := credentials.NewClientTLSFromCert(pool, "")
perRPC, _ := oauth.NewServiceAccountFromFile("service-account.json", scope)
conn, _ := grpc.Dial(
	"greeter.googleapis.com",
	grpc.WithTransportCredentials(creds),
	grpc.WithPerRPCCredentials(perRPC),
)
// error handling omitted
client := pb.NewGreeterClient(conn)
// ...
```
