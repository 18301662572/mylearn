package token

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

//实现grpc SomeMethod方法，首先进行 Authentication认证，认证通过了才可以继续操作

type grpcServer struct{ auth *Authentication }

func (p *grpcServer) SomeMethod(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	if err := p.auth.Auth(ctx); err != nil {
		return nil, err
	}
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

//认证
func (a *Authentication) Auth(ctx context.Context) error {

	//首先通过metadata.FromIncomingContext从ctx上下文中获取元信息，
	// 然后取出相应的认证信息进行认证。如果认证失败，则返回一个codes.Unauthenticated类型地错误。
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var appid string
	var appkey string
	if val, ok := md["user"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}
	if appid != a.User || appkey != a.Password {
		return grpc.Errorf(codes.Unauthenticated, "invalid token")
	}
	return nil
}
