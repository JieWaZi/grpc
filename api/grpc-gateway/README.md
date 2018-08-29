# 使用grpc-gateway并与Consul进行集成
---

### Handler
由于使用了grpc-gateway的框架，所以使用protoc命令编译后就会生成.pb.go文件，我们可以查看生成的文件中的代码,我截取了主要部分的代码。

``` golang
// RegisterFirstCalculationServiceHandlerFromEndpoint is same as RegisterFirstCalculationServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterFirstCalculationServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterFirstCalculationServiceHandler(ctx, mux, conn)
}

// RegisterFirstCalculationServiceHandler registers the http handlers for service FirstCalculationService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterFirstCalculationServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterFirstCalculationServiceHandlerClient(ctx, mux, NewFirstCalculationServiceClient(conn))
}

// RegisterFirstCalculationServiceHandler registers the http handlers for service FirstCalculationService to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "FirstCalculationServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "FirstCalculationServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "FirstCalculationServiceClient" to call the correct interceptors.
func RegisterFirstCalculationServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client FirstCalculationServiceClient) error {

	mux.Handle("POST", pattern_FirstCalculationService_AddCalculation_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		req.Header.Set("Grpc-Metadata-scope", "")
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FirstCalculationService_AddCalculation_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FirstCalculationService_AddCalculation_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}
```
以上代码包含三个方法：
- RegisterFirstCalculationServiceHandlerFromEndpoint
- RegisterFirstCalculationServiceHandler
- RegisterFirstCalculationServiceHandlerClient

看主要方法XXXHandlerClient：
可以看出该方法就是将Handler注册到ServerMux（路由管理器）中，而ServerMux本身也实现了Handler接口的ServeHTTP方法，所以可得出只需要将所有生成的Handler都注册到ServerMux即可实现HTTP请求的应答。  
而XXXHandler及XXXHandlerFromEndpoint则是都是调用上面的方法，并进行了进一步封装。
