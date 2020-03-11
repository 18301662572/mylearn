package main

//数据库:用户表，书籍表，书籍分类表
//书籍id（一个是自增的，另一个是雪花算法 snowflake生成唯一ID）
//用户登录页面+书籍列表页面
//gin框架
//Cookie+ginSession
//logrus 第三方日志库
//logagent一个监控组件：监控服务器信息（网卡IO，内存，磁盘，cpu等）根据IP发送到kafka, 从kafka取值，通过grafana展示出来
// 						(使用Kafka、Elasticsearch、Grafana搭建业务监控系统)[https://blog.csdn.net/tonywu1992/article/details/83506671]
//logagent启动， 1.收集日志信息   2.收集系统信息  （一样的流程）
