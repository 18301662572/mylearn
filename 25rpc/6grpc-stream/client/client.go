package main

import (
	"code.oldbody.com/studygolang/mylearn/25rpc/6grpc-stream"
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

//GRPC流 客户端实现流服务
//完成了完整的流接收和发送支持。

var client pb.HelloServiceClient

func main() {
	//客户端需要先调用Channel方法获取返回的流对象
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//在客户端我们将发送和接收操作放到两个独立的Goroutine。

	// 首先是向服务端发送数据
	go func() {
		for {
			if err := stream.Send(&pb.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	//在循环中接收服务端返回的数据
	go func() {
		for {
			reply, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			fmt.Println(reply.GetValue())
		}
	}()
}
