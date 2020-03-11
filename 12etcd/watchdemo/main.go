package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//etcd watch 实时监控etcd里面指定key的变化

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.42.133:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v\n", err)
		return
	}
	defer cli.Close()

	//watch
	watchCh := cli.Watch(context.Background(), "s4")
	for wresp := range watchCh {
		for _, evt := range wresp.Events {
			fmt.Printf("type:%s Key:%s Value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value) //evt.Kv.Key=s4
		}
	}
}
