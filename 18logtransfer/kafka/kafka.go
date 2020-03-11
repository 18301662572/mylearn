package kafka

import (
	"code.oldbody.com/studygolang/mylearn/logtransfer/es"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

//1.初始化kafka连接
//2.从kafka里取出日志数据

func Init(address []string, topic string) (err error) {
	// 创建新的消费者
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		logrus.Errorf("fail to start consumer, err:%v\n", err)
		return
	}
	// 拿到指定topic下面的所有分区列表
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		logrus.Errorf("fail to get list of partition:err%v\n", err)
		return
	}
	logrus.Info(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest) //从最新的开始读，一直读到的是最新的
		if err != nil {
			logrus.Errorf("failed to start consumer for partition %d,err:%v\n",
				partition, err)
			continue
		}
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				logrus.Infof("Partition:%d Offset:%d Key:%s Value:%s",
					msg.Partition, msg.Offset, msg.Key, msg.Value)
				es.PutLogData(msg) //为了将同步流程异步化，所以将取出的日志数据先放到channel中
				var m1 map[string]interface{}
				err = json.Unmarshal(msg.Value, &m1)
				if err != nil {
					logrus.Warningf("unmarshal msg failed,err:%v", err)
					continue
				}
				es.PutLogData(m1)
			}
		}(pc)
	}
	return
}
