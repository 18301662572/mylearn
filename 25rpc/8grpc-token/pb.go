package token

//grpc请求,响应结构体

type HelloRequest struct {
	Name string
}

type HelloReply struct {
	Message string
}
