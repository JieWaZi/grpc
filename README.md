
# golang使用grpc和grpc-gateway

### 参考文档
> [grpc中文文档](http://doc.oschina.net/grpc?t=58008)  
[protobuf官方文档](https://developers.google.com/protocol-buffers/)    
[protobuf中文文档](https://blog.csdn.net/u011518120/article/details/54604615)    
[grpc-gateway官方文档](https://github.com/grpc-ecosystem/grpc-gateway)  
### grpc介绍


### 安装Protobuf

#### 介绍两种安装方式：

一、下载压缩文件

[下载地址](https://github.com/protocolbuffers/protobuf/releases) 

由于本人是os系统，所以选择的是protoc-3.6.1-osx-x86_64.zip。解压后在bin文件夹下会有一个编译好的protoc可执行文件。配置路径,/Users/xxx/protoc/bin为我解压portoc文件的路径

``` shell
vim ~/.bash_profile

export PROTOC=/Users/xxx/protoc/bin
export PATH=$PATH:$PROTOC

source ~/.bash_profile
```

或则直接将可执行文件放入`$GOPATH`下的bin文件中，并配置一下`$GOBIN`

在进行测试是否安装成功

``` shell
protoc --version
```
返回相关版本信息即成功,如下显示

``` shell
libprotoc 3.6.0
```

二、下载源码自行编译

``` shell
git clone https://github.com/google/protobuf

cd protobuf

./autogen.sh

./configure

make

make check

sudo make install
```

### 下载golang插件

由于都是github的项目，最好是可以科学上网
``` shell
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
go get -u golang.org/x/text
go get -u golang.org/x/net
go get -u golang.org/x/tools
```
使用该命令可能无法安装

``` shell
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
git clone https://github.com/golang/tools.git $GOPATH/src/golang.org/x/tools
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

cd $GOPATH/src/

go install google.golang.org/grpc
```

编译protoc-gen-go

``` shell
cd $GOPATH/github.com/golang/protobuf/protoc-gen-go
go build
```
编译好的protoc-gen-go插件会`$GOPATH/github.com/golang/protobuf/protoc-gen-go`目录下，然后将生成的二进制可执行protoc-gen-go文件放置于`$GOPATH/bin`下，以便protoc能够找到protoc-gen-go。

### 简单的HelloWorld项目

#### 编写proto文件

``` proto
syntax = "proto3";

option go_package = "proto";

package helloworld;

service Message {
    rpc SendMessage (SendMessageRequest) returns (SendMessageResponse) {}
}

message SendMessageRequest {
    string toWho = 1;
    string message = 2;
}

message SendMessageResponse {
    string fromWho = 1;
    string message = 2;
}
```

#### 生成*.pb.go文件
```
protoc --go_out=plugins=grpc:. *.proto
```

#### 编写server端
``` go
package main

// server.go

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/grpc/proto"
)

type grpcServer struct{}

func (s *grpcServer) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	return &pb.SendMessageResponse{FromWho: in.ToWho, Message: "receive message : " + in.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageServer(s, &grpcServer{})
	s.Serve(lis)
}

```

#### 编写client端

``` go
package main

//client.go

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/grpc/proto"
)
func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessageClient(conn)

	r, err := c.SendMessage(context.Background(), &pb.SendMessageRequest{ToWho: "wjj",Message:"helloWorld"})
	if err != nil {
		log.Fatal("could not SendMessage: %v", err)
	}
	log.Printf("SendMessage: %s", r.Message)
}
```

分别开启server和client端口，显示结果如下：
```
 SendMessage: receive message : helloWorld
 ```


### 使用grpc-gateway

#### 获取grpc-gateway
``` shell

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```
分别进入protoc-gen-grpc-gateway,protoc-gen-swagger及protoc-gen-go文件夹下将编译好的二进制文件放到$GOPATH/bin中，使其相关命令生效。

#### 编写proto文件

``` proto
syntax = "proto3";

option go_package = "proto";

import "google/api/annotations.proto";

message SimpleMessaegeRequest {
    string message = 1;
}

message SimpleMessaegeResponse {
    string name = 1;
    int32 age = 2;
}


service MessageService {

    rpc SayHello (SimpleMessaegeRequest) returns (SimpleMessaegeResponse) {
        option (google.api.http) = {
           post: "/test/hello"
           body: "*"
        };
    };
}

```

其中 google/api/annotations.proto文件在grpc-gateway/third_party/googleapis/google/api/文件夹下
将该文件夹放到项目中。

#### 编译proto文件
``` shell
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  path/to/your_service.proto
```

#### 生成反向代理
``` shell
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  path/to/your_service.proto
```

#### 编写server
``` go
package main

import (
	"context"
	pb "github.com/grpc/proto"
	"net"
	"log"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, request *pb.SimpleMessaegeRequest) (*pb.SimpleMessaegeResponse, error) {
	return &pb.SimpleMessaegeResponse{
		Name: request.Message,
		Age:  10,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9100")
	if err != nil {
		log.Fatal("some wrong")
	}
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("some wrong")
	}
}

```

#### 编写proxy
``` go
package main

import (
  "flag"
  "net/http"

  "github.com/golang/glog"
  "golang.org/x/net/context"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"
	
  gw "path/to/your_service_package"
)

var (
  echoEndpoint = flag.String("echo_endpoint", "localhost:9100", "endpoint of YourService")
)

func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
  if err != nil {
    return err
  }

  return http.ListenAndServe(":8080", mux)
}

func main() {
  flag.Parse()
  defer glog.Flush()

  if err := run(); err != nil {
    glog.Fatal(err)
  }
}
```
分别运行proxy和server。

使用postman发起post请求，url为localhost:8080/test/hello。