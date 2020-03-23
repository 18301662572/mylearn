package main

import (
	pb "code.oldbody.com/studygolang/mylearn/25rpc/12grpc-rest"
	"context"
	"fmt"
)

//GRPC服务

//Post 返回值的数据结构
type ReqParams struct {
	Value string
}

type RestService struct {
}

func (r *RestService) Get(c context.Context, req *pb.StringMessage) (resp *pb.StringMessage, err error) {
	resp = new(pb.StringMessage)
	resp.Value = fmt.Sprintf("Get %s", req.Value)
	return
}
func (r *RestService) Post(c context.Context, req *pb.StringMessage) (resp *pb.StringMessage, err error) {
	resp = new(pb.StringMessage)
	resp.Value = fmt.Sprintf("Post:%s", req.Value)
	return
}
