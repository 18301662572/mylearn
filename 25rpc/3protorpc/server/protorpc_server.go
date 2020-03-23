package main

import (
	pb "code.oldbody.com/studygolang/mylearn/25rpc/3protorpc"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

//protobuf来定义RPC方法及其请求响应参数，并使用第三方的protorpc库来生成RPC服务注册代码。
//protobuf server  "google.golang.org/grpc"

// 算术运算结构体
type ArithService struct {
}

var a = ArithService{}

// 乘法运算方法
func (a *ArithService) Multiply(ctx context.Context, req *pb.ArithRequest) (res *pb.ArithResponse, err error) {
	res = new(pb.ArithResponse) //初始化
	res.Pro = req.A * req.B
	return
}

// 除法运算方法
func (a *ArithService) Divide(ctx context.Context, req *pb.ArithRequest) (res *pb.ArithResponse, err error) {
	if req.B == 0 {
		return nil, errors.New("divide by zero")
	}
	res = new(pb.ArithResponse) //初始化
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return
}

func main() {
	//1.监听
	lis, err := net.Listen("tcp", "127.0.0.1:8097")
	if err != nil {
		fmt.Printf("监听异常：%s \n", err)
		return
	}
	fmt.Printf("开始监听：%s \n", "127.0.0.1:8097")
	//2、实例化gRPC
	s := grpc.NewServer()
	//3.在gRPC上注册微服务
	//第二个参数接口类型的变量
	pb.RegisterArithServiceServer(s, &a)
	//注册反射
	reflection.Register(s)
	//4.启动gRPC服务端
	s.Serve(lis)

}
