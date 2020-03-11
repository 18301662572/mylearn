package main

import (
	//"net/rpc"
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

//声明请求的参数结构体
type ArithRequest struct {
	A, B int
}

//声明一个响应客户端参数的结构体
type ArithResponse struct {
	//乘积
	Pro int
	//商
	Quo int
	//余数
	Rem int
}

//1.golang中实现RPC-调用服务
//func main() {
//	//连接远程rpc
//	rp, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
//	if err != nil {
//		log.Fatal(err)
//	}
//	req := ArithRequest{14, 6}
//	var resp ArithResponse
//	err = rp.Call("Arith.Multiply", req, &resp)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%d * %d = %d \n", req.A, req.B, resp.Pro)
//	err = rp.Call("Arith.Divide", req, &resp)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%d / %d, 商= %d，余数=%d \n", req.A, req.B, resp.Quo, resp.Rem)
//}

//2.jsonRPC调用
func main() {
	//连接远程rpc
	rp, err := jsonrpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	req := ArithRequest{14, 6}
	var resp ArithResponse
	err = rp.Call("Arith.Multiply", req, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d * %d = %d \n", req.A, req.B, resp.Pro)
	err = rp.Call("Arith.Divide", req, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d / %d, 商= %d，余数=%d \n", req.A, req.B, resp.Quo, resp.Rem)
}
