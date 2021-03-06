package main

import (
	pb "code.oldbody.com/studygolang/mylearn/19my-micro/gRPC/proto"
	"context"
	"github.com/golang/go/src/pkg/fmt"
	"google.golang.org/grpc"
	"net"
)

//定义服务端实现约定的接口
type UserInfoService struct {
}

var u = UserInfoService{}

//实现服务端需要实现的接口
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name
	//在数据库查用户信息
	if name == "zs" {
		resp = &pb.UserResponse{
			Id:   1,
			Name: name,
			Age:  22,
			//切片字段
			Hobby: []string{"Sing", "Run", "Read"},
		}
	}
	err = nil
	return
}

func main() {
	//1.监听
	addr := "127.0.0.1:8080"
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
	pb.RegisterUserInfoServiceServer(s, &u)
	//4.启动gRPC服务端
	s.Serve(lis)
}
