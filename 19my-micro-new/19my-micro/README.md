# rpc




## go-micro

1. .proto文件生成 （.pb.go文件 和  .micro.go文件）的命令
```text
protoc -I . --micro_out=. --go_out=. ./hello.proto
（. 生成的文件放在平级目录下）
protoc -I . --micro_out=../src/share/pb --go_out=../src/share/pb ./user.proto
（../src/share/pb生成的文件放在上级src/share/pb目录下）	
```
