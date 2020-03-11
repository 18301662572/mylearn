package main

import (
	pb "code.oldbody.com/studygolang/mylearn/19my-micro/go-micro02/proto"
	context "context"
	"fmt"
	"github.com/go-micro"
	"github.com/go-micro/errors"
	"log"
)

//go.micro.api.example 服务端实现
//启动服务：1.打开cmd，输入命令 micro api --handler=rpc，给rpc开启一个http api 8080监听端口，方便http访问
//		   2.右键运行
//Client访问：使用http进行访问（Postman工具）
//POST http://localhost:8080/example/call  Body:JSON,传入name参数
//POST http://localhost:8080/example/foo/bar

type Example struct{}

type Foo struct{}

func (e *Example) Call(ctx context.Context, req *pb.CallRequest, resp *pb.CallResponse) error {
	log.Print("收到 Example.Call 请求")
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}
	resp.Message = "RPC Call 收到了你的请求" + req.Name
	return nil
}

func (f *Foo) Bar(ctx context.Context, req *pb.EmptyRequest, resp *pb.EmptyResponse) error {
	log.Print("收到 Foo.Bar 请求")
	return nil
}

func main() {
	//1.得到微服务实例
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)
	//2.初始化
	service.Init()
	//3.服务注册
	//注册Example接口
	err := pb.RegisterExampleHandler(service.Server(), new(Example))
	if err != nil {
		fmt.Println(err)
		return
	}
	//注册Foo接口
	err = pb.RegisterFooHandler(service.Server(), new(Foo))
	if err != nil {
		fmt.Println(err)
		return
	}
	//4.启动微服务
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
