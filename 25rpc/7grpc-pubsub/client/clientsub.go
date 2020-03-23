package client

import (
	"code.oldbody.com/studygolang/mylearn/25rpc/7grpc-pubsub"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

//客户端进行 订阅信息
func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 构造客户端的发布订阅对象
	client := pb.NewPubsubServiceClient(conn)
	//客户端订阅信息
	stream, err := client.Subscribe(context.Background(), &pb.String{Value: "golang:"}) //.SubscribeTopic(context.Background(), &pb.String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}
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
}
