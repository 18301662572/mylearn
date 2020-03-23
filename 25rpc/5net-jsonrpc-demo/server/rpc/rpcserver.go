package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//支持TCP协议的 jsonrpc server
//在TCP协议之上运行jsonrpc服务，并且通过nc命令行工具成功实现了RPC方法调用
//Linux nc命令：nc -l 1234

//直接向架设了RPC服务的TCP服务器发送json数据模拟RPC方法调用：
//echo -e '{"method":"HelloService.Hello","params":["hello"],"id":1}' | nc localhost 1234
//返回结果
//{"id":1,"result":"hello:hello","error":null}
//其中id对应输入的id参数，result为返回的结果，error部分在出问题时表示错误信息。对于顺序调用来说，
//id不是必须的。但是Go语言的RPC框架支持异步调用，当返回结果的顺序和调用的顺序不一致时，可以通过id来识别对应的调用。

//因此无论采用何种语言，只要遵循同样的json结构，以同样的流程就可以和Go语言编写的RPC服务进行通信。
//这样我们就实现了跨语言的RPC。

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
