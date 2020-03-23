package main

import (
	pb "code.oldbody.com/studygolang/mylearn/25rpc/12grpc-rest"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

//grpc网关 -> HTTP反向代理服务器的入口点

//请求rest服务
//$ curl localhost:8080/get/gopher
//{"value":"Get: gopher"}
//$ curl localhost:8080/post -X POST --data '{"value":"grpc"}'
//{"value":"Post: grpc"}

// 首先通过runtime.NewServeMux()函数创建路由处理器，
// 然后通过RegisterRestServiceHandlerFromEndpoint函数将RestService服务相关的REST接口中转到后面的GRPC服务。
// grpc-gateway提供的runtime.ServeMux类也实现了http.Handler接口，因此可以和标准库中的相关函数配合使用。

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	//创建路由处理器
	mux := runtime.NewServeMux()

	//创建一个[]grpc.DialOption
	opts := []grpc.DialOption{grpc.WithInsecure()}

	//RestService服务相关的REST接口中转到后面的GRPC服务
	err := pb.RegisterRestServiceHandlerFromEndpoint(
		ctx, mux, "localhost:1234",
		opts,
	)
	if err != nil {
		log.Fatal(err)
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(":8080", mux)
}
