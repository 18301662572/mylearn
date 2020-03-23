package filter

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"log"
)

//GRPC 截取器 支持链式
//开源的grpc-ecosystem项目中的go-grpc-middleware包已经基于GRPC对截取器实现了链式截取器的支持。

func main() {
	myServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		//filter1, filter2, ...
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		//filter1, filter2, ...
		)),
	)
	log.Print(myServer)
}
