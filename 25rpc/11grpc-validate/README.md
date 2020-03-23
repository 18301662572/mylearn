# GRPC Validate 验证器
https://github.com/gogo/protobuf<br>
https://hant.kutu66.com//GitHub/article_137710<br>
通过生成的验证函数，并结合GRPC的截取器，我们可以很容易为每个方法的输入参数和返回值进行验证。<br>

# 注：
```text
go get XXXX时，
需要先设置：export PATH=${PATH}:${GOPATH}/bin
才能在GoPath/bin目录下生成相应的.exe文件
```

```text
syntax = "proto2";第二版的Protobuf有个默认值特性，可以为字符串或数值类型的成员定义默认值。
syntax = "proto3";在第三版的Protobuf中不再支持默认值特性，但是我们可以通过扩展选项自己模拟默认值特性。
```
### 安装和使用：
```text
github.com/mwitkow/go-proto-validators 已经基于Protobuf的扩展特性实现了功能较为强大的验证器功能。
安装和使用：
    1.protoc 编译器希望在执行 $PATH 上找到名为 proto-gen-XYZ（protoc-gen-govalidators）的插件。 所以首先
        export PATH=${PATH}:${GOPATH}/bin
    2.go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
    这样在gopath/bin下会生成 protoc-gen-govalidators.exe 插件，之后使用插件生成.go验证文件
```
### 生成validate.validator.pb.go的命令：
```text
protoc  --proto_path=${GOPATH}/src  --proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf  --proto_path=.  --gogo_out=.  --govalidators_out=gogoimport=true:.  validate.proto
```
### Install protoc-gen-gogo:
```text
Install protoc-gen-gogo:
go get github.com/gogo/protobuf
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/jsonpb
go get github.com/gogo/protobuf/protoc-gen-gogo
go get github.com/gogo/protobuf/gogoproto
```

