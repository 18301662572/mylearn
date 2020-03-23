```text
logAgent-日志收集系统架构图（Log Agent -> Kafka ->log transfer->ES->Kibana->浏览器）
sysAgent-系统信息收集架构图 （Sys Agent（gopsutil包收集服务器性能信息）->Kafka->sys transfer->influxDB->Grafana->浏览器） 
```

### 1.**log Agent流程梳理**

详情图片见“images/*” （包含日志收集系统总的架构图）

```go
1.加载配置文件

2.1初始化kafka生产者producer (*sarama)

2.2初始化存放日志数据的channel（msgChan），将读日志和发送日志改为异步执行

2.3后台起一个goroutine，从msgChan中接收日志数据发送到kafka

3.1初始化一个连接etcd的client（clientv3）

4.去etcd中根据给定的key获取配置文件，配置文件是一个切片，切片中的每一项是日志文件的路径和topic

5.启动一个后台的goroutine，监听etcd中的日志收集项是否发生变化

6.1创建一个全局的tailTask管理者

6.2遍历传过来的配置，每有一个配置项就启动一个tailTask，启动一个后台的goroutine去执行日志收集

6.3启动后台的goroutine去等新的配置来，从一个无缓冲的channel中接收新配置
```

### 2.**FileBeat 学习链接：**
[https://github.com/elastic/logstash-forwarder]
<br/>

### **3.日志收集项目总结**
```go
logagent

​	etcd:  put \ set \ watch  类似的项目: consul

​	tail包用法

​	kafka: 组件、内部的原理 (zookeeper: 跟etcd一样，目前使用率不如etcd)

​sysagent  ：第十五天的系统信息

​	gopsutl包：类似的现成的监控`open-falcon`

​	influxDB：时序数据库 ，类似的项目还有`openTSDB`

​	grafana:  数据可视化

 logtransfer

​	Elasticsearch (ES): 存储+全文检索，搜索引擎，各个概念和内部的组件构成

​	Kibana:可视化

​注：ELK-->  ES+ Logstash + Kibana
```


