# GRPC REST接口
相关网址：<br/>
详解： https://www.jishuchi.com/read/GO/696<br/>
参考：https://segmentfault.com/a/1190000008106582?utm_source=tuicool&utm_medium=referral<br/>
grpc-gateway网关网址 ：https://github.com/grpc-ecosystem/grpc-gateway<br/>


GRPC服务一般用于集群内部通信，如果需要对外暴露服务一般会提供等价的REST接口。<br/>
通过REST接口比较方便前端JavaScript和后端交互。开源社区中的grpc-gateway项目就实现了<br/>
将GRPC服务转为REST服务的能力。<br/>


### grpc-gateway的工作原理
(见图 images/grpc-gateway.png)
```text
通过在Protobuf文件中添加路由相关的元信息，
通过自定义的代码插件生成路由相关的处理代码，最终将REST请求转给更后端的GRPC服务处理。
```

### 步骤
```text
1.创建一个.proto路由文件
2.依赖导入"google/api/annotations.proto"; （需要翻墙或者将包下载到本地）
    下载到本地：
    google/api/annotations.proto
    google/api/http.proto
    google/api/httpbody.proto
    google/protobuf/any.proto
    google/protobuf/descriptor.proto
3.安装protoc , protoc-gen-grpc-gateway插件
    https://github.com/grpc-ecosystem/grpc-gateway
    export PATH=${PATH}:${GOPATH}/bin (.exe文件加入$GOPATH/bin目录)
    go install \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
        github.com/golang/protobuf/protoc-gen-go
4.生成服务端 .router.pb.go文件 ，生成网关（gateway）router.pb.gw.go文件 ， 生成接口规范 router.swagger.json文件
5.编写grpc服务端业务逻辑代码及注册服务。（grpc-server）
6.编写grpc网关代码,即对外REST服务（grpc-gateway）
7.当GRPC和REST服务全部启动之后，就可以用curl请求REST服务了：
```

    
### 下载包网址：
```text
Google API
网址：https：//github.com/google/googleapis
protobuf
网址：https://github.com/golang/protobuf

google/api/包: https://github.com/googleapis/googleapis/blob/master/google/api<br/>
google/protobuf/descriptor包: https://github.com/golang/protobuf/blob/master/protoc-gen-go/descriptor/descriptor.proto<br/>
google/protobuf/any.proto包: https://github.com/golang/protobuf/tree/master/ptypes/any<br/>
```

### 生成命令(windows)
```text
//.router.pb.go
//生成服务端 
 protoc -I C:/protoc/protoc/bin -I. \
        -I$GOPATH/src \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --go_out=plugins=grpc:. \
        router.proto

//router.pb.gw.go
//生成网关使用部分
protoc -I C:/protoc/protoc/bin -I. \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --grpc-gateway_out=logtostderr=true:. \
       router.proto

//router.swagger.json
在对外公布REST接口时，我们一般还会提供一个Swagger格式的文件用于描述这个接口规范。
在网页中提供REST接口的文档和测试等功能。
protoc -I. \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --swagger_out=. \
    router.proto
```

### curl 请求TEST服务
```text
$ curl localhost:8080/get/gopher
{"value":"Get: gopher"}
$ curl localhost:8080/post -X POST --data '{"value":"grpc"}'
{"value":"Post: grpc"}
```

### 问题
```text
解决网址：https://www.cnblogs.com/zjhblogs/p/11505432.html
GIT_TAG="v1.2.0" # change as needed
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go
```
