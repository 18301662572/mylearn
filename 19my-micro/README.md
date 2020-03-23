# RPC :远程过程调用（Remote Procedure Call）,是一个计算机通信协议
该协议允许运行于一台计算机的程序调用另一台计算机的子程序，而程序员无需额外的为这个交互作用编程。

##### golang的RPC必须符合4个条件
```text
1.结构体字段首字母要大写，要跨域访问，所以大写
2.函数名必须首字母大写（可以序列化导出的）
3.函数第一个参数是接受参数，第二个参数是返回给客户端的参数（第二个参数必须是指针类型）
4.函数必须有一个返回值error
```

**注：golang官方的net/rpc库使用 encoding/gob进行编解码，支持tcp和http数据传输方式，由于其他语言不支持gob编解码方式，所以golang的RPC只支持golang开发的服务器与客户端之间的交互**<br/>
**另外，官方提供 net/rpc/jsonrpc 库实现RPC方法，jsonrpc通过json格式编解码，支持跨语言调用,目前jsonrpc库是基于tcp协议实现的，暂不支持http传输方式**<br/>


## grpc

1.生成go文件<br>
  打开命令行，输入命令生成接口文件(user.pb.go) 
```text
protoc -I . --go_out=plugins=grpc:. ./user.proto
```

2.go-micro框架，.proto文件生成  .pb.go文件 和  .micro.go文件的命令：	
```text
 protoc -I . --micro_out=. --go_out=. ./hello.proto
 (. 生成的文件放在平级目录下）
 protoc -I . --micro_out=../src/share/pb --go_out=../src/share/pb ./user.proto
 (../src/share/pb生成的文件放在上级src/share/pb目录下）
 注： ./当前目录   ../上级目录
```	



​							
