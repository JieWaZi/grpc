package main

import (
	"flag"
	"net/http"
	"golang.org/x/net/context"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"github.com/grpc/api/grpc-gateway/http/handler"
	"log"
	"github.com/grpc/api/consul"
	_ "github.com/grpc/api/grpc-gateway/server"
	"fmt"
)

var (
	reg  = flag.String("server", "127.0.0.1:8500", "consul address")
	host = flag.String("proxy_host", "127.0.0.1", "register host")
	port = flag.String("proxy_port", "8989", "register port")
)

func run() error {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := newGateway(ctx)

	return http.ListenAndServe(fmt.Sprintf("%s:%s", *host, *port), mux)
}

// 将HTTP请求的服务转为GRPC，即使用GRPC生成的Handler
func newGateway(ctx context.Context) http.Handler {
	mux := runtime.NewServeMux()
	for _, sh := range handler.ListHTTPHandlers() {

		// 新建一个地址解析
		r := consul.NewResolver(sh.Service)
		b := grpc.RoundRobin(r)

		dialOptions := []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithBalancer(b),
		}
		// 于consul进行连接通信，找到对应的服务地址并建立连接
		conn, err := grpc.Dial(*reg, dialOptions...)
		if err != nil {
			log.Fatal("dial grpc failed", "error", err, "service", sh.Service)
		}
		err = sh.RegisterHttpHandler(ctx, mux, conn)
		if err != nil {
			log.Print("register service handler failed", "error", err, "service", sh.Service)
		}
	}

	return mux
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal("Failed to start rpc proxy server", "error", err)
	}
}
