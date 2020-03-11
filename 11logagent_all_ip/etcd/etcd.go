package etcd

import (
	"code.oldbody.com/studygolang/mylearn/11logagentallip/common"
	"code.oldbody.com/studygolang/mylearn/11logagentallip/tailfile"
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//etcd相关操作

var client *clientv3.Client

func Init(address []string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		return
	}
	return
}

//GetConf 拉取日志收集配置项的函数
func GetConf(key string) (collectEntryList []common.CollectEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		log.Errorf("get conf from etcd failed,err:%v", err)
		return
	}
	if len(resp.Kvs) == 0 {
		log.Warn("get len:0 conf from etcd")
		return
	}
	ret := resp.Kvs[0]
	//ret.Value//json 格式字符串
	//value:[{"path":"g:/logs/s4.log","topic":"test_log"},{"path":"g:/logs/s2.log","topic":"web_log"}]
	fmt.Println(ret.Value)
	err = json.Unmarshal(ret.Value, &collectEntryList)
	if err != nil {
		log.Errorf("json unmarshal failed,err:%v", err)
		return
	}
	return
}

//WatchConf 监控etcd中日志收集项配置 变化 的函数
func WatchConf(key string) {
	for {
		watchCh := client.Watch(context.Background(), key)
		for wresp := range watchCh {
			log.Info("get new conf from etcd")
			for _, evt := range wresp.Events {
				fmt.Printf("type:%s Key:%s Value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value) //evt.Kv.Key=key
				var newConf []common.CollectEntry
				if evt.Type == clientv3.EventTypeDelete {
					//如果是删除
					log.Warnf("FBI warning:etcd delete the key: %v!!!", evt.Kv.Key)
					tailfile.SendNewConf(newConf) //没有任何接受就是阻塞的
					continue
				}
				//type:PUT Key:collect_log_conf Value:[{"path":"g:/logs/s4.log","topic":"test_log"},{"path":"g:/logs/s2.log","topic":"web_log"},{"path":"g:/logs/new.log","topic":"new_log"},{"path":"g:/logs/new1.log","topic":"new1_log"}]
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					log.Errorf("json unmarshal new conf failed,err:%v", err)
					continue
				}
				//告诉tailfile模块应该启用新的配置了！
				tailfile.SendNewConf(newConf) //没有任何接受就是阻塞的
			}
		}
	} //for循环，让他一直监控变化
}
