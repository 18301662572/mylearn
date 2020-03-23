package filter

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

//GRPC 截取器针对方法提供的, 单个 (截取器也非常适合前面对Token认证工作)
//GRPC中的grpc.UnaryInterceptor和grpc.StreamInterceptor分别对普通方法和流方法提供了截取器的支持。
//不过GRPC框架中只能为每个服务设置一个截取器，因此所有的截取工作只能在一个函数中完成。
//开源的grpc-ecosystem项目中的go-grpc-middleware包已经基于GRPC对截取器实现了链式截取器的支持。(main.go)

// 函数的ctx和req参数就是每个普通的RPC方法的前两个参数。
// 第三个info参数表示当前是对应的那个GRPC方法，
// 第四个handler参数对应当前的GRPC方法函数。
// 首先是日志输出info参数，然后调用handler对应的GRPC方法函数。
func filter1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("fileter:", info)
	return handler(ctx, req)
}

//截取器增加了对GRPC方法异常的捕获
func filter2(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("fileter:", info)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return handler(ctx, req)
}

func main() {
	//使用filter截取器函数，只需要在启动GRPC服务时作为参数输入即可
	grpc.NewServer(grpc.UnaryInterceptor(filter1))
	grpc.NewServer(grpc.UnaryInterceptor(filter2))
}
