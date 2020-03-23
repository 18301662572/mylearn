# go-mico

go-micro项目架构
```text
demo
proto           .proto原文件
src
    api-srv:    微服务网关
    user-srv:   用户服务
        db      实例化数据库，表操作
        entity  数据库表到实体的映射
        handler 业务逻辑
        main    go micro 写微服务的流程
                   //1.得到微服务实例
                   //2.初始化
                   //3.服务注册（将实现接口的业务实体注册到微服务里）
                   //4.启动微服务
    share:      配置，工具类，生成的proto文件
    vendor:     存放第三方的库
    dbhepler:   数据库相关常见表
    api-srv:    微服务的网关，统一处理请求
```

```text
1.Go Micro（可插拔的插件化架构）**是微服务的框架**，是一个插件化的基础框架，基于此可以构建微服务。<br>
在架构之外，他默认实现了 consul 作为服务发现，通过protobuf 和 json进行编解码
2.主要功能：
    服务发现：自动服务注册和名称解析。服务发现是微服务开发的核心。 
    负载均衡：基于服务发现构建的客户端负载均衡。 
    消息编码：基于内容类型的动态消息编码。 
    请求响应：基于 RPC 的请求/响应，支持双向流。 
    Async Messaging(异步通信)：PubSub 是异步通信和事件驱动架构的一流公民。 事件通知是微服务开发的核心模式 
    可插拔接口：Go Micro 为每个分布式系统抽象使用 Go 接口 
    注：插件地址： https://github.com/micro/go-plugins
3.通信流程：
    Server端监听，Client端调用，Brocker将信息推送过来进行处理。Register服务注册和发现。
4.核心接口：
    go-micro 之所以可以高度订制和他的框架结构是分不开的， 由 8 个主要的inteface 构成了 go-micro 的框架结构 ，
每一个 interface 都可以根据自己的需求重新实现。
    见图“images/”
5.go micro接口详解（8个接口）
    1.Transort通信接口：服务发送和接收的最终实现。
    2.Codec 编码接口：编解码方式，默认实现方式是protobuf
    3.Registry注册接口：服务的注册和发现。目前实现的有consul,mdns,etcd,zookeeper,kubernets等方式
    4.Selector负载均衡：Selector是客户端级别的负载均衡，当有客户端向服务发送请求时，selector根据不同的算法
从Regietery中的主机列表得到可用的Service节点进行通信，目前实现的有循环算法和随机算法，默认是随机算法。	
    5.Broker发布订阅接口：消息发布和订阅的接口
    6.Client客户端接口：请求服务的接口，他封装Transort和Codec进行rpc调用，也封装了Broker进行信息发布。
    7.Server服务端接口： 监听等待 rpc 请求， 监听 broker 的订阅信息，等待信息队列的推送等 
    8.Service接口：是 Client 和 Server 的封装，他包含了一系列的方法使用初始值去初始化Service 和 Client，
使我们可以很简单的创建一个 rpc 服务
6.go micro实践
    1.go-micro 安装
        查看的网址： https://github.com/micro/
        cmd 中输入下面 3 条命令下载， 会自动下载相关的很多包
        go get -u -v github.com/micro/micro
        go get -u -v github.com/micro/go-micro
        go get -u -v github.com/micro/protoc-gen-micro 
    2..proto文件生成  .pb.go文件 和  .micro.go文件的命令：		
        protoc -I . --micro_out=. --go_out=. ./hello.proto
        （. 生成的文件放在平级目录下）
        protoc -I . --micro_out=../src/share/pb --go_out=../src/share/pb ./user.proto
        （../src/share/pb生成的文件放在上级src/share/pb目录下）
         注： ./当前目录   ../上级目录
    3.go micro 写微服务的流程：
        //1.得到微服务实例
        //2.初始化
        //3.服务注册（将实现接口的业务实体注册到微服务里）
        //4.启动微服务
            //1.得到微服务实例
            service := micro.NewService(
                //设置微服务的名字，用来做访问用的，
                micro.Name("hello"),
            )
            //2.初始化
            service.Init()
            //3.服务注册（将实现接口的业务实体注册到微服务里）
            err := pb.RegisterHelloHandler(service.Server(), new(Hello))
            if err != nil {
                fmt.Println(err)
                return
            }
            //4.启动微服务
            if err = service.Run(); err != nil {
                log.Fatal(err)
            }

7.微服务网关 （api-srv）
    1.编写请求处理 
  ​		Access-Control-Allow-Origin 设置允许的跨域地址
  ​		Access-Control-Allow-Methods POST, GET 设置跨越请求允许的请求方式
  ​		Access-Control-Allow-Headers Content-Type, Content-Length, Accept-Encoding,
  X-Token, X-Client 设置跨越请求允许的数据格式
        Access-Control-Allow-Credentials true 设置跨越请求是否可携带证书 
        //RPC请求,跨域请求
            if origin := r.Header.Get("Origin"); true {
                w.Header().Set("Access-Control-Allow-Origin", origin)
            }
            w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding,X-Token, X-Client")
            w.Header().Set("Access-Control-Allow-Credentials", "true")

8.项目的访问测试
    使用consul 作服务发现 (安装consul.exe,配置环境变量)
    1.打开cmd，输入：consul agent -dev 监听服务，可以很容易看出从哪个端口去监听
    2.打开consul控制台页面：
        http://localhost:8500/ui/dc1/services
    3.启动一个微服务，右键启动user服务
    4.consul控制台页面 Services列表 可以看到 启动的微服务
        http-访问微服务
        1.启动服务网关--(api-srv)
            右键启动 api-srv 服务 （设置端口号：8888）
        2.Postman 测试访问
            Postman访问 http://localhost:8888/user/userService/InsertUser  +JSON参数
```


