package main

import (
	pb "code.oldbody.com/studygolang/mylearn/25rpc/3protorpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

//protobuf来定义RPC方法及其请求响应参数，并使用第三方的protorpc库来生成RPC服务注册代码。
//protobuf client

func main() {
	//1.创建与gRPC服务端的连接
	//grpc.WithInsecure() 建立一个安全连接（跳过了对服务器证书的验证）
	conn, err := grpc.Dial("127.0.0.1:8097", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Printf("连接异常：err %s\n", err)
		return
	}
	//2.实例化gRPC客户端
	client := pb.NewArithServiceClient(conn)
	//3.组装参数
	req := new(pb.ArithRequest)
	req.A = 9
	req.B = 2
	//4.调用接口
	resp, err := client.Multiply(context.Background(), req)
	if err != nil {
		fmt.Printf("响应异常：arith error: %s\n", err)
		return
	}
	fmt.Printf("%d * %d = %d\n", req.GetA(), req.GetB(), resp.GetPro())

	resp, err = client.Divide(context.Background(), req)
	if err != nil {
		log.Fatalln("arith error ", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, resp.Quo, resp.Rem)
}
