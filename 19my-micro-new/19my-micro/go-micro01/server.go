package main

import (
	pb "code.oldbody.com/studygolang/mylearn/19my-micro/go-micro01/proto"
	context "context"
	"fmt"
	"github.com/go-micro"
	"log"
)

//go-micro 服务端server实现
//启动服务：1.打开cmd，输入命令 micro api --handler=rpc，给rpc开启一个http api 监听端口，方便http访问
//		   2. 右键运行
//Client访问:使用cmd命令行进行访问
//命令： micro call hello Hello.Info {\"userName\":\"zs\"}

//声明结构体
type Hello struct {
}

//实现接口方法
func (h *Hello) Info(ctx context.Context, req *pb.InfoRequest, resp *pb.InfoResponse) error {
	resp.Msg = "你好" + req.UserName
	return nil
}

func main() {
	//1.得到微服务实例
	service := micro.NewService(
		//设置微服务的名字，用来做访问用的，
		micro.Name("hello"),
	)
	//2.初始化
	service.Init()
	//3.服务注册
	err := pb.RegisterHelloHandler(service.Server(), new(Hello))
	if err != nil {
		fmt.Println(err)
		return
	}
	//4.启动微服务
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
