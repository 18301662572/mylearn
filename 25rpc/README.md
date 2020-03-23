# golang实现RPC的几种方式

RPC详解:  https://www.jishuchi.com/read/GO/691<br/>
RPC实现简介： https://studygolang.com/articles/14336<br/>
PROTOBUF快速上手指南：<br>
https://idoubi.cc/2017/12/02/protobuf%E5%BF%AB%E9%80%9F%E4%B8%8A%E6%89%8B%E6%8C%87%E5%8D%97/<br/>
GRPCURL：https://github.com/fullstorydev/grpcurl<br/>

# 注：
```text
go get XXXX时，
需要先设置：export PATH=${PATH}:${GOPATH}/bin
才能在GoPath/bin目录下生成相应的.exe文件
```

```text
1.golang官方的net/rpc库使用encoding/gob进行编解码，支持tcp或http数据传输方式，
由于其他语言不支持gob编解码方式，所以使用net/rpc库实现的RPC方法没办法进行跨语言调用。
2.golang官方还提供了net/rpc/jsonrpc库实现RPC方法，JSON RPC采用JSON进行数据编解码，
因而支持跨语言调用。但目前的jsonrpc库是基于tcp协议实现的，也可以在http协议上提供jsonrpc服务
3.除了golang官方提供的rpc库，还有许多第三方库为在golang中实现RPC提供支持，
大部分第三方rpc库的实现都是使用protobuf进行数据编解码，根据protobuf声明文件自动生成rpc方法定义与服务注册代码，
在golang中可以很方便的进行rpc服务调用。
JSON和protobuf是支持多语言的，所以使用jsonrpc和protorpc实现的RPC方法我们是可以在其他语言中进行调用的。
```
### 项目架构
```text
1net-rpc                 RPC使用gob编解码,不能跨语言调用,支持tcp或http数据传输方式
    rpc-server           服务端
    rpc-client           客户端
2net-rpcjsonrpc          JSON RPC采用JSON进行数据编解码，因而支持跨语言调用，使用json传输数据编解码性能不高等,基于tcp协议，也可以在http协议上提供jsonrpc服务
    jsonrpc-server       服务端
    jsonrpc-client       客户端
3protorpc                protobuf定义RPC方法及其请求响应参数，并使用第三方的protorpc库来生成RPC服务注册代码。支持跨语言调用,比json编码性能好
    arith.proto          protobuf文件
    protorpc_server      服务端
    protorpc_client      客户端
4net-rpc-demo            对RPC的client及server进行了封装调用的 实践
5net-jsonrpc-demo        对jsonrpc的 实践 (TCP / HTTP 协议上提供jsonrpc服务)
6grpc-stream             GRPC流
7grpc-pubsub             GRPC发布与订阅/GRPC简单实现了一个跨网络的发布和订阅服务
8grpc-token              GRPC token认证 （对每个grpc方法进行权限认证）/（CS根证书是对每个grpc链接认证）
9grpc-filter             GRPC截取器
10grpc-web               GRPC与Web服务共存
11grpc-validate          GRPC验证器
12grpc-rest              GRPC REST接口 通过grpc-gateway实现对内RPC对外REST
13pbgo                   基于Protobuf的框架
```

### ProtoBuf->GRPC ：Protobuf定义语言无关的RPC服务接口
```text
1.首先，需要安装protobuf及protoc可执行命令，可以参考此篇文章：protobuf快速上手指南，生成go文件
2.要先安装protorpc库：
    go get google.golang.org/grpc
3.bin目录下载好工具包： 
    protoc.exe，protoc-gen-go.exe，protoc-gen-protorpc.exe
4.然后使用protoc工具生成代码：
    protoc -I . --go_out=plugins=grpc:. ./arith.proto
    或者 protoc --go_out=. arith.proto
    （其中go_out参数告知protoc编译器去加载对应的protoc-gen-go工具，
    然后通过该工具生成代码，生成代码放到当前目录。最后是一系列要处理的protobuf文件的列表。）
    执行protoc命令后，在与arith.proto文件同级的目录下生成了一个arith.pb.go文件，里面包含了RPC方法定义和服务注册的代码。
5.go-mic框架需要工具包：
    protoc-gen-micro.exe 
6.下载protoc工具：（同3）
    https://github.com/google/protobuf/releases
7.安装针对Go语言的代码生成插件（同2）
    go get github.com/golang/protobuf/protoc-gen-go
8.在protoc-gen-go内部已经集成了一个叫grpc的插件，可以针对grpc生成代码
    protoc --go_out=plugins=grpc:. hello.proto
9.GRPC构建在HTTP/2协议之上;
```

### GRPC调试工具--grpcurl/grpcui
```text
grpcurl（go语言） 和 grpcui（js语言） 都是调试grpc的利器，前者用于命令行，类似curl工具；后者是以web的形式进行调试的，类似postman工具。
有了这两款工具，我们不用写任何客户端代码，也能方便的调试接口数据。
这两款工具的作者是同一人：http://github.com/fullstorydev 。
因为网络原因grpcui下载不下来，没办法安装。
grpcurl
1.install grpcurl
    go get github.com/fullstorydev/grpcurl
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl
    因为网络原因，go install不成功！所以
    grpcurl.exe 下载目录：
    https://github.com/fullstorydev/grpcurl/releases
2.grpcurl.exe 放在$GOPATH/bin目录下，或者添加环境变量
3.grpcurl操作详情：
    https://www.jishuchi.com/read/GO/ch4-rpc-ch4-08-grpcurl.md
4.grpccurl操作总结：
a.server服务端添加反射代码
    //注册反射
    reflection.Register(s)      
b.注：本地不能使用 127.0.0.1；只能localhost 表示
      -plaintext: 用来忽略tls证书的验证过程。
      参数名：遵循.pb.go文件的请求结构json tag （因为curl以json的格式传递参数）
c.grpcurl调用方法：
    grpcurl -plaintext -d "{\"a\":8,\"b\":2}" localhost:8097  pb.ArithService/multiply
d.grpcurl用list命令查看服务列表    
    grpcurl -plaintext localhost:8097 list
e.grpcurl用list命令查看服务的方法列表     
    grpcurl -plaintext localhost:8097 list pb.ArithService
f.grpcurl用describe子命令查看更详细的描述信息     
    grpcurl -plaintext localhost:8097 describe pb.ArithService
g.grpcurl用describe子命令查看参数类型的信息： 
    grpcurl -plaintext localhost:8097 describe pb.ArithService.divide
```

### 目前存在的问题
```text
1.没有启动TLS加密 -->8grpc-token/CA根证书
客户端与服务器建立连接，使用grpc.WithInsecure()选项跳过了对服务器证书的验证，
没有启用证书的GRPC服务在和客户端进行的是明文通讯，信息面临被任何第三方监听的风险。
为了保障GRPC通信不被第三方监听篡改或伪造，我们可以对服务器启动TLS加密特性。
    CA根证书：使用命令给服务器和客户端分别生成私钥和证书，
             通过一个安全可靠的根证书分别对服务器和客户端的证书进行签名
2.grpc是从缓存读取数据吗？
3.grpc支持http协议吗？
    首先GRPC是建立在HTTP/2版本之上，如果HTTP不是HTTP/2协议则必然无法提供GRPC支持。
    同时，每个GRPC调用请求的Content-Type类型会被标注为”application/grpc”类型。
4.protobuf与msgpack哪个更好？
```

### GRPC 认证
```text

证书认证 
1.CA证书：基于证书的认证是针对每个GRPC链接的认证。
实现了一个服务器和客户端进行双向证书验证的通信可靠的GRPC系统。
原因：
    GRPC建立在HTTP/2协议之上，对TLS提供了很好的支持。我们前面章节中GRPC的服务都没有提供证书支持，
    因此客户端在链接服务器中通过grpc.WithInsecure()选项跳过了对服务器证书的验证。
    没有启用证书的GRPC服务在和客户端进行的是明文通讯，信息面临被任何第三方监听的风险。
    为了保障GRPC通信不被第三方监听篡改或伪造，我们可以对服务器启动TLS加密特性。
CA根证书: （服务器端证书文件+签名)
    通过一个安全可靠的根证书分别对服务器和客户端的证书进行签名。
    这样客户端或服务器在收到对方的证书后可以通过根证书进行验证证书的有效性。
2.Token认证：为每个GRPC方法调用提供了认证支持，这样就基于用户Token对不同的方法访问进行权限管理。
3.截取器： GRPC中的grpc.UnaryInterceptor和grpc.StreamInterceptor分别对普通方法和流方法提供了截取器的支持。
          截取器也非常适合前面对Token认证工作
```

### GRPC和Web服务共存
```text
GRPC构建在HTTP/2协议之上，因此我们可以将GRPC服务和普通的Web服务架设在同一个端口之上。
因为目前Go语言版本的GRPC实现还不够完善，只有启用了TLS协议之后才能将GRPC和Web服务运行在同一个端口。
```