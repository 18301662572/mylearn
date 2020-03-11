##influxdb 时序数据库##
[https://www.liwenzhou.com/posts/Go/go_influxdb/]

## 安装

官网下载页面(https://portal.influxdata.com/downloads/)

**Windows**
下载链接:https://dl.influxdata.com/influxdb/releases/influxdb-1.7.7_windows_amd64.zip

**Mac**
下载链接:https://dl.influxdata.com/influxdb/releases/influxdb-1.7.7_darwin_amd64.tar.gz

或者:
```go
brew update
brew install influxdb
```

### 基本命令
官方文档：[https://docs.influxdata.com/influxdb/v1.7/introduction/getting-started/]

### 介绍
influxdb时序数据库，开源，Go开发的，
集群功能是收费的
类似项目：OpenTSDB (集群不收费)

### 操作
见“images/*”
注：插入数据操作，记得tags和fileds的区别