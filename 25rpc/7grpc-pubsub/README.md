# GRPC实现发布和订阅模式

### 架构
```text
client            
         clientpub      客户端发布数据
         clientsub      客户端订阅数据
server 
         server         服务端实现发布订阅功能
pubsub.pb.go            protoc使用grpc插件实现的.go文件        
pubsub.proto            .proto文件
docker-demo             docker项目中实现的一个pubsub的demo
```

### 注
```text
protobuf文件中的stream流可以生成 流 Send(),Reav()方法。
在发布和订阅模式中，由调用者主动发起的发布行为类似一个普通函数调用，
而被动的订阅者则类似GRPC客户端单向流中的接收者。
```

### 步骤
```text
基于GRPC简单实现了一个跨网络的发布和订阅服务：
1.pubsub.proto 创建Publish(),Subscribe()(stream String)方法
2.protoc --go_out=plugins=grpc:. hello.protopubsub
生成.pb.go 
```
```go
//pubsub.pb.go文件
//服务端接口
type PubsubServiceServer interface {
    Publish(context.Context, *String) (*String, error)
    Subscribe(*String, PubsubService_SubscribeServer) error
}
//客户端接口
type PubsubServiceClient interface {
    Publish(context.Context, *String, ...grpc.CallOption) (*String, error)
    Subscribe(context.Context, *String, ...grpc.CallOption) (
        PubsubService_SubscribeClient, error,
    )
}
//stream流生成的 Send()方法
type PubsubService_SubscribeServer interface {
    Send(*String) error
    grpc.ServerStream
}
````
```text
3.服务端
    1.创建服务端发布订阅结构体
    2.构造服务端的发布订阅对象
    3.实现发布方法
    4.实现订阅方法
4.客户端
    1.构造客户端的发布订阅对象
    2.从客户端向服务器发布信息  ---> 调用服务端的发布方法
    3.客户端订阅服务器信息      ---> 调用客户端的订阅方法  
```
