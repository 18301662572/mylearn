package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

//将日志数据写入Elasticsearch

type ESClient struct {
	client      *elastic.Client  //es连接
	index       string           //es中连接那个数据库
	logDataChan chan interface{} //接受日志的channel
}

var (
	esClient *ESClient
)

func Init(address, index string, goroutine_num, maxSize int) (err error) {
	client, err := elastic.NewClient(elastic.SetURL("http://" + address))
	if err != nil {
		// Handle error
		panic(err)
	}
	logrus.Info("connect to es success")
	esClient = &ESClient{
		client:      client,
		index:       index,
		logDataChan: make(chan interface{}, maxSize),
	}
	//从channel中取数据,写入到kafka中
	for i := 0; i < goroutine_num; i++ {
		go sendToES()
	}

	return
}

func sendToES() {
	for m1 := range esClient.logDataChan {
		put1, err := esClient.client.Index(). //获取所有的数据库
							Index(esClient.index). //Index索引的数据库
							BodyJson(m1).
							Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		logrus.Infof("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	}
}

//通过一个首字母大写的函数，从包外接受msg,放到chan中
func PutLogData(msg interface{}) {
	esClient.logDataChan <- msg
}
