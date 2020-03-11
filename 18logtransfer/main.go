package main

import (
	"code.oldbody.com/studygolang/mylearn/logtransfer/es"
	"code.oldbody.com/studygolang/mylearn/logtransfer/kafka"
	"code.oldbody.com/studygolang/mylearn/logtransfer/model"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

//log transfer
//从 kafka 消费日志数据，写入ES

func main() {
	var cfg = new(model.Config)
	//0.读配置文件 `go-ini`包
	err := ini.MapTo(cfg, "./config/logtransfer.ini")
	if err != nil {
		log.Errorf("load config failed,err:%v", err)
		panic(err)
	}
	log.Info("load config success!")
	//2.连接ES
	err = es.Init(cfg.ESConf.Address, cfg.ESConf.Index, cfg.ESConf.GoNum, cfg.ESConf.MaxSize)
	if err != nil {
		log.Errorf("init es failed,err:%v", err)
		panic(err)
	}
	log.Info("init es success!")
	//3.连接kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		log.Errorf("init kafka failed,err:%v", err)
		panic(err)
	}
	log.Info("init kafka success!")
	//在这儿停顿！
	select {}
}
