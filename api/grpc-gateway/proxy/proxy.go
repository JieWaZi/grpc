package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"

	gw "github.com/grpc/proto"
)

var (
	echoEndPoint  = flag.String("FirstCalculationService", "localhost:9100", "endpoint of YourService")
	echoEndPoint1 = flag.String("SecondCalculationService", "localhost:9200", "endpoint of YourService")
)

func run1() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterFirstCalculationServiceHandlerFromEndpoint(ctx, mux, *echoEndPoint, opts)
	if err != nil {
		return err
	}
	err = gw.RegisterSecondCalculationServiceHandlerFromEndpoint(ctx, mux, *echoEndPoint1, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8989", mux)

}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run1(); err != nil {
		glog.Fatal(err)
	}
}
