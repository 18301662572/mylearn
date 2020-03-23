package token

import (
	"google.golang.org/grpc"
	"log"
)

//Token认证

//CA根证书： 针对每个GRPC 链接 的认证。
//Token： 针对每个GRPC 方法 的认证，基于用户Token对不同的方法访问进行权限管理。

func main() {
	auth := Authentication{
		User:     "gopher",
		Password: "password",
	}
	//通过grpc.WithPerRPCCredentials函数将Authentication对象转为grpc.Dial参数。
	// 因为这里没有启用安全链接，需要传人grpc.WithInsecure()表示忽略证书认证
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//...
}
