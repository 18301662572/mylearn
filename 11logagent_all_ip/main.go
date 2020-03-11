package main

import (
	"code.oldbody.com/studygolang/mylearn/11logagentallip/common"
	"code.oldbody.com/studygolang/mylearn/11logagentallip/etcd"
	"code.oldbody.com/studygolang/mylearn/11logagentallip/kafka"
	"code.oldbody.com/studygolang/mylearn/11logagentallip/tailfile"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

//日志收集的客户端(全版+加入了获取本机ip)
//类似的开源项目还有filebeat

//目标：如果有多台服务器需要收集日志，添加各个服务器ip版本
//1. conf.ini配置中 collect_key=collect_log_%s_conf （%s：每个服务器ip）
//2. common中获取本机ip,
//3. 将etcd获取日志收集配置项的key中的%s替换成本机ip

type KafkaConfig struct {
	Address  string `ini:"address"`
	Topic    string `ini:"topic"`
	ChanSize int64  `ini:"chan_size"`
}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}

type EtcdConfig struct {
	Address    string `ini:"address"`
	CollectKey string `ini:"collect_key"`
}

//整个logagent的配置
type Config struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
	EtcdConfig    `ini:"etcd"`
}

func run() {
	select {}
}

func main() {
	//-1:获取本机ip,为后续etcd拉取要收集日志的配置项使用
	ip, err := common.GetOutboundIP()
	if err != nil {
		log.Errorf("get ip failed,err:%v", err)
		return
	}

	var configObj = new(Config)
	//0.读配置文件 `go-ini`包
	err = ini.MapTo(configObj, "./conf/config.ini")
	if err != nil {
		log.Errorf("ini load config failed,err:%v", err)
		return
	}
	kafkaAddr := configObj.KafkaConfig.Address
	//1.初始化 连接kafka
	err = kafka.Init([]string{kafkaAddr}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		log.Errorf("init kafka failed,err:%v", err)
		return
	}
	log.Infof("init kafka success!")

	//2.初始化etcd连接
	err = etcd.Init([]string{configObj.EtcdConfig.Address})
	if err != nil {
		log.Errorf("init etcd failed,err:%v", err)
		return
	}
	//3.从etcd中拉取要收集日志的配置项
	collectKey := fmt.Sprintf(configObj.EtcdConfig.CollectKey, ip)
	allConf, err := etcd.GetConf(collectKey)
	if err != nil {
		log.Errorf("get conf from etcd failed,err:%v", err)
		return
	}
	fmt.Println(allConf)

	//后台goroutine,不能阻塞，一直监控（ 派一个小弟去监控etcd中 configObj.EtcdConfig.CollectKey 对应值的变化 ）
	go etcd.WatchConf(collectKey)

	//4.根据配置中的日志路径初始化tail（程序第一次加载使用）
	err = tailfile.Init(allConf) //把从etcd获取的配置项传到Init中
	if err != nil {
		log.Errorf("init tail failed,err:%\v", err)
		return
	}
	log.Info("init tail success!")

	//4.把日志通过sarama发往kafka,在tailfile文件中实现

	//让程序一直处于阻塞状态
	run()
}
