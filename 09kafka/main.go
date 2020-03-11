package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"math/rand"
	"strconv"
)

//基于sarama第三方库开发的 kafka client,往kafka里面发送消息

func main() {
	//1.生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //Ack 设置生产者发送完数据是否需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //partition 新选择一个分区partition
	config.Producer.Return.Successes = true                   //确认 成功交付的消息将在success channel返回

	//2.连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.42.133:9092"}, config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return
	}
	defer client.Close()

	//3.封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log" //shopping
	msg.Value = sarama.StringEncoder("this is a test log" + strconv.Itoa(rand.Intn(9)))

	//4.发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed,err:", err)
		return
	}
	fmt.Printf("分区pid:%v 偏移offset:%v\n", pid, offset)
}
