package handler

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

//RpcToHttpHandlerRegister is the microservice entry point which register RPC service handler
type HandlerRegister func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

var httpHandlers []HTTPHandler

// HTTPHandler is used to expose RPC services through HTTP/HTTPS
type HTTPHandler struct {
	Service             string
	RegisterHttpHandler HandlerRegister
}

// RegisterHTTPHandler is used by the RPC services to register itself so that it can expose RPC through HTTP
func RegisterHTTPHandler(service string, hf HandlerRegister) {
	httpHandlers = append(httpHandlers, HTTPHandler{
		Service:             service,
		RegisterHttpHandler: hf,
	})
}

// ListHTTPHandlers returns all Handlers which translate the HTTP requests into RPC requests
func ListHTTPHandlers() []HTTPHandler {
	return httpHandlers
}
