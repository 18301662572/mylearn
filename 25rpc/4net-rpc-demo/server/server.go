package server

import (
	inter "code.oldbody.com/studygolang/mylearn/25rpc/4net-rpc-demo/interface"
	"log"
	"net"
	"net/rpc"
)

//基于RPC接口规范编写真实的服务端代码：

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	inter.RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
