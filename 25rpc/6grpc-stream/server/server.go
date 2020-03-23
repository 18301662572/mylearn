package main

import (
	"code.oldbody.com/studygolang/mylearn/25rpc/6grpc-stream"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
)

//GRPC流 服务端实现流服务

//RPC是远程函数调用，因此每次调用的函数参数和返回值不能太大，否则将严重影响每次调用的响应时间。
// 因此传统的RPC方法调用对于上传和下载较大数据量场景并不适合。
// 同时传统RPC模式也不适用于对时间不确定的订阅和发布模式。
// 为此，GRPC框架针对服务器端和客户端分别提供了流特性。

//服务端在循环中接收客户端发来的数据，如果遇到io.EOF表示客户端流被关闭，如果函数退出表示服务端流关闭。
// 生成返回的数据通过流发送给客户端，双向流数据的发送和接收都是完全独立的行为。
// 需要注意的是，发送和接收的操作并不需要一一对应，用户可以根据真实场景进行组织代码。

type HelloServiceImpl struct {
}

func (p *HelloServiceImpl) Hello(context.Context, *pb.String) (*pb.String, error) {
	return &pb.String{Value: ""}, nil
}

func (p *HelloServiceImpl) Channel(stream pb.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		reply := &pb.String{Value: "hello:" + args.GetValue()}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

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
	pb.RegisterHelloServiceServer(s, &HelloServiceImpl{})
	//4.启动gRPC服务端
	s.Serve(lis)
}
