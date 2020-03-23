package main

import (
	pb "code.oldbody.com/studygolang/mylearn/25rpc/12grpc-rest"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

//注册rpc服务
func main() {
	//1.监听
	addr := "127.0.0.1:1234"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常：%s \n", err)
		return
	}
	fmt.Printf("开始监听：%s \n", addr)
	//2、实例化gRPC
	s := grpc.NewServer()
	//3.在gRPC上注册微服务
	//第二个参数要接口类型的变量
	var rest = new(RestService)
	pb.RegisterRestServiceServer(s, rest)
	//4.启动gRPC服务端
	s.Serve(lis)
}
