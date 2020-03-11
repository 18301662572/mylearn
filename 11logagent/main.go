package main

import (
	"code.oldbody.com/studygolang/mylearn/logagent/kafka"
	"code.oldbody.com/studygolang/mylearn/logagent/tailfile"
	"fmt"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"strings"
	"time"
)

//日志收集的客户端
//类似的开源项目还有filebeat

//目标：收集指定目录下的日志文件，发送到kafka中
//1.往kafka发数据 sarama
//2.使用tail包读日志文件

type KafkaConfig struct {
	Address  string `ini:"address"`
	Topic    string `ini:"topic"`
	ChanSize int64  `ini:"chan_size"`
}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}

//整个logagent的配置
type Config struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
}

//真正的业务逻辑
func run() (err error) {
	//从logfile通过TailObj取出log，通过client发送到kafka
	for {
		//循环读数据
		line, ok := <-tailfile.TailObj.Lines //chan tail.Line
		if !ok {
			log.Warnf("tail file close reopen,filename:%s\n", tailfile.TailObj.Filename)
			time.Sleep(time.Second) //读取出错，等一秒
			continue
		}
		log.Infof("msg:", line.Text)
		// 去除换行符
		strLineText := strings.Replace(line.Text, "\r", "", -1)
		if len(strLineText) == 0 {
			log.Info("出现空行，直接跳过...")
			continue
		}
		//利用通道将同步的代码改为异步的
		//把读出来的一行日志包装成kafka里面的msg类型，丢到通道中
		msg := &sarama.ProducerMessage{}
		msg.Topic = "web_log"
		msg.Value = sarama.StringEncoder(strLineText)
		fmt.Println(msg.Value)
		//丢到通道中
		kafka.ToMsgChan(msg)
	}
	return
}

func main() {
	var configObj = new(Config)
	//0.读配置文件 `go-ini`包
	err := ini.MapTo(configObj, "./conf/config.ini")
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
	//2.根据配置中的日志路径初始化tail
	err = tailfile.Init(configObj.CollectConfig.LogFilePath)
	if err != nil {
		log.Errorf("init tail %s failed,err:%\v", configObj.CollectConfig.LogFilePath, err)
		return
	}
	log.Infof("init tail %s success!", configObj.CollectConfig.LogFilePath)
	//3.把日志通过sarama发往kafka
	err = run()
	if err != nil {
		log.Errorf("run failed,err:%v", err)
		return
	}
}
