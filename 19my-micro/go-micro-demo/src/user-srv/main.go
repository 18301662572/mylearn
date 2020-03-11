package main

import (
	"code.oldbody.com/studygolang/mylearn/19my-micro/go-micro-demo/src/share/config"
	"code.oldbody.com/studygolang/mylearn/19my-micro/go-micro-demo/src/share/pb"
	"code.oldbody.com/studygolang/mylearn/19my-micro/go-micro-demo/src/share/utils/log"
	"code.oldbody.com/studygolang/mylearn/19my-micro/go-micro-demo/src/user-srv/db"
	"code.oldbody.com/studygolang/mylearn/19my-micro/go-micro-demo/src/user-srv/handler"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

//程序的入口
//server服务的实现

func main() {
	log.Init("user")
	logger := log.Instance()
	//1.创建service
	service := micro.NewService(
		micro.Name(config.Namespace+"user"),
		micro.Version("latest"),
	)
	//2.初始化service
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("user-srv服务运行时的打印...")
			//初始化db
			db.Init(config.MysqlDNS)
			//3.注册服务
			err := pb.RegisterUserServiceHandler(service.Server(),
				handler.NewUserHandler(), server.InternalHandler(true))
			if err != nil {
				fmt.Println(err)
			}
		}),
		//定义服务停止后做的事情
		micro.AfterStop(func() error {
			logger.Info("user-srv服务停止后的打印...")
			return nil
		}),
		micro.AfterStart(func() error {
			logger.Info("user-srv服务启动前的打印...")
			return nil
		}),
	)
	logger.Info("启动user-srv服务...")
	//4.启动service
	if err := service.Run(); err != nil {
		//logger.Panic("user-srv服务启动失败...")
		logger.Info("user-srv服务启动失败...")
	}
}
