package client

import (
	"code.oldbody.com/studygolang/mylearn/25rpc/7grpc-pubsub"
	"context"
	"google.golang.org/grpc"
	"log"
)

//从客户端向服务器 发布信息
func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 构造客户端的发布订阅对象
	client := pb.NewPubsubServiceClient(conn)
	//客户端发布信息
	_, err = client.Publish(context.Background(), &pb.String{Value: "golang: hello Go"})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(context.Background(), &pb.String{Value: "docker: hello Docker"})
	if err != nil {
		log.Fatal(err)
	}
}
