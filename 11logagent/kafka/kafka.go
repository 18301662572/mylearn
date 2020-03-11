package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

//kafka相关操作

var (
	client  sarama.SyncProducer
	msgChan chan *sarama.ProducerMessage
)

//Init 是初始化全局的kafka连接
func Init(address []string, chanSize int64) (err error) {
	//1.生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //Ack 设置生产者发送完数据是否需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //partition 新选择一个分区partition
	config.Producer.Return.Successes = true                   //确认 成功交付的消息将在success channel返回

	//2.连接kafka
	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		return err
	}
	//初始化 MsgChan
	msgChan = make(chan *sarama.ProducerMessage, chanSize)
	//起一个后台的goroutine从MsgChan中读数据
	go sendMsg()
	return
}

//从MsgChan中读取msg,发送给kafka
func sendMsg() {
	for {
		select {
		case msg := <-msgChan:
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				logrus.Error("send msg to kafka failed,err:", err)
				return
			}
			logrus.Info("send msg to  kafka success! 分区pid:%v 偏移offset:%v", pid, offset)
		}
	}
}

func ToMsgChan(msg *sarama.ProducerMessage){
	msgChan <- msg
}
