package server

import (
	"code.oldbody.com/studygolang/mylearn/25rpc/7grpc-pubsub"
	"context"
	"github.com/moby/moby/pkg/pubsub"
	"strings"
	"time"
)

//基于GRPC的流特性构造一个发布和订阅系统。

//在发布和订阅模式中，由调用者主动发起的发布行为类似一个普通函数调用，
//而被动的订阅者则类似GRPC客户端单向流中的接收者。

//PubsubService 服务端发布订阅结构体
type PubsubService struct {
	pub *pubsub.Publisher
}

//NewPubsubService 构造服务端的发布订阅对象
func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

//实现发布方法
func (p *PubsubService) Publish(ctx context.Context, arg *pb.String) (*pb.String, error) {
	p.pub.Publish(arg.GetValue())
	return &pb.String{}, nil
}

//实现订阅方法
func (p *PubsubService) Subscribe(arg *pb.String, stream pb.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&pb.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}
