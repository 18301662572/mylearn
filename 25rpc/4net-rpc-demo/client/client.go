package client

import (
	inter "code.oldbody.com/studygolang/mylearn/25rpc/4net-rpc-demo/interface"
	"log"
)

//rpc客户端

func main() {
	client, err := inter.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
}
