package main

import (
	"code.oldbody.com/studygolang/mylearn/logagentnew/etcd"
	"code.oldbody.com/studygolang/mylearn/logagentnew/kafka"
	"code.oldbody.com/studygolang/mylearn/logagentnew/tailfile"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

//日志收集的客户端 (全版)
//类似的开源项目还有filebeat

//目标1：(第一版实现)
//1.收集指定目录下的日志文件，发送到kafka中
//2.往kafka发数据 sarama
//3.使用tail包读日志文件

//目标2：
//1.之前的文件的路径是从config.ini中的collect节点获取到的，是固定的且只能写一个
//更改后：
//1.将文件路径的配置项写入到etcd中（F:\Learn\Go\GOWork\src\code.oldbody.com\studygolang\mylearn\12etcd\demo\main 中的test()方法 测试文件：1.exe 2.exe 3.exe），
//  从etcd中获取文件路径。 etcd有实时监控某个key并返回给消费者的优势，不需要重启配置就可以一直修改配置项；
//2.从etcd中获取key-value集合（json格式）,获取到每个文件的路径，循环集合获取每个文件，将文件的 path，topic,文件内容写入到tailTask，循环输出每行文件内容
//3.保证程序一直处于阻塞状态

//目标3：etcd一直监控 日志文件配置项 (Key对应的Value)的变化，处理修改后的配置项
//1. 后台起一个goroutine， 使用etcd watch（）实时监控collect_log_conf 对应的Value的变化，将变化后的Value (json反序列化 json.Unmarshal()）传入tailfile的confChan通道中。 goroutine一直处于阻塞状态，一直在等Value变化。
//2. tailfile的confChan 接收到配置项,对应新的配置项对之前运行的tailTask进行管理
//3. tail新建一个tailTaskMgr 管理类，包含tailTaskMap（目前所有的tailTask任务），collectEntryList（所有配置项），confChan（tailfile中等待新配置的通道）
//4. 将tailfile类中的逻辑代码及confChan通道合并到tailTaskMgr中，在tailTaskMgr中对接收到的新配置进行管理
//5. 创建一个后台goroutine对新配置进行管理 tailTaskMgr.watch()
//	 循环判断 confChan（[]common.CollectEntry）
//	  5.1. 如果没有发生改变的配置项，不做操作
//    5.2  新添加的配置项，要创建新的tailTask,并开启该对应的goroutine
//    5.3  已经删除的配置项，要停掉对应的goroutine (通过调用tailTaskMgr中的ctx.cancel()，其对应的goroutine中的ctx.Done()可以收到值，通过 <-ctx.Done(): retrun 将该goroutine退出)，同时一定要在集合中删除对应的配置项
//注意： watch() 一定要for{} ,不能只读一次 etcd.watch(),ttMgr.watch(),tailfile.run()

//暂留的问题：
//如果logagent停了需要记录上一次的位置，参考filebeat

//总结：日志文件配置项目 （go-ini->etcd->tail->kafka(sarama,zookeeper)）
//1.在 config.ini 配置文件中，配置连接地址，etcd监控的key（key的Value格式（[{"path":"g:/logs/s4.log","topic":"test_log"},{"path":"g:/logs/s2.log","topic":"web_log"}]））
//2.通过(go-ini包)获取conf.ini获取连接地址和etcd key 的值
//3.使用 etcd 监控文件配置项的信息（Key对应的Value）
//4.获取到配置项，循环配置项，将每个配置项交由 tail包 循环读取每一个该文件每一行的数据，将每行数据跟topic打包成sarama识别的类型，发送到kafka
//注：必须用go mod ,因为sarama 使用v1.19.0版本

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

	//2.初始化etcd连接
	err = etcd.Init([]string{configObj.EtcdConfig.Address})
	if err != nil {
		log.Errorf("init etcd failed,err:%v", err)
		return
	}
	//3.从etcd中拉取要收集日志的配置项
	allConf, err := etcd.GetConf(configObj.EtcdConfig.CollectKey)
	if err != nil {
		log.Errorf("get conf from etcd failed,err:%v", err)
		return
	}
	fmt.Println(allConf)

	//后台goroutine,不能阻塞，一直监控（ 派一个小弟去监控etcd中 configObj.EtcdConfig.CollectKey 对应值的变化 ）
	go etcd.WatchConf(configObj.EtcdConfig.CollectKey)

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
