package main

import (
	hello_pb "github.com/chai2010/pbgo/examples/hello.pb"
	"log"
	"net/http"
)

type HelloService struct{}

func (p *HelloService) Hello(request *hello_pb.String, reply *hello_pb.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func (p *HelloService) Echo(in *hello_pb.Message, out *hello_pb.Message) error {
	return nil
}
func (p *HelloService) Static(in *hello_pb.String, out *hello_pb.StaticFile) error {
	return nil
}

//测试REST服务：
//curl localhost:8080/hello/vgo

func main() {
	router := hello_pb.HelloServiceHandler(new(HelloService))
	log.Fatal(http.ListenAndServe(":8080", router))
}
