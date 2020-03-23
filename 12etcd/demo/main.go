package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

//代码连接etcd
//重命名exe: go build -o rename.exe

//http访问 （修改key）：
//curl http://192.168.42.138:2379/v2/keys/collect_log_192.168.42.133_conf
// -XPUT -d value="[{"path":"g:/logs/s4.log","topic":"test_log"}]"

//{
//"action": "set",
//"node": {
//"createdIndex": 1,
//"key": "/collect_log_192.168.42.133_conf",
//"modifiedIndex": 1,
//"value": "[{"path":"g:/logs/s4.log","topic":"test_log"},{"path":"g:/logs/s2.log","topic":"web_log"},{"path":"g:/logs/new1.log","topic":"new1_log"}]"
//}
//}


func bodymain() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.42.138:2379"}, // Linux服务器 // []string{"192.168.42.133:2379"}, //Windows服务器
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v\n", err)
		return
	}
	defer cli.Close()

	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "beibei", "beibei is a test")
	if err != nil {
		fmt.Printf("put to etcd failed,err:%v\n", err)
		return
	}
	cancel()

	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gr, err := cli.Get(ctx, "beibei")
	if err != nil {
		fmt.Printf("get from etcd failed,err:%v\n", err)
		return
	}
	for _, ev := range gr.Kvs {
		fmt.Printf("key:%s value:%s\n", ev.Key, ev.Value)
	}
	cancel()
}

func test() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.42.133:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v\n", err)
		return
	}
	defer cli.Close()

	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//strJson := `[{"path":"g:/logs/s4.log","topic":"test_log"},{"path":"g:/logs/s2.log","topic":"web_log"}]`
	//strJson := `[{"path":"g:/logs/s4.log","topic":"test_log"},{"path":"g:/logs/s2.log","topic":"web_log"},{"path":"g:/logs/new.log","topic":"new_log"}]`
	strJson := `[{"path":"g:/logs/s4.log","topic":"test_log"},{"path":"g:/logs/s2.log","topic":"web_log"},{"path":"g:/logs/new1.log","topic":"new1_log"}]`
	//_, err = cli.Put(ctx, "collect_log_conf", strJson)
	_, err = cli.Put(ctx, "collect_log_192.168.42.133_conf", strJson)
	if err != nil {
		fmt.Printf("put to etcd failed,err:%v\n", err)
		return
	}
	cancel()

	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//gr, err := cli.Get(ctx, "collect_log_conf")
	gr, err := cli.Get(ctx, "collect_log_192.168.42.133_conf")
	if err != nil {
		fmt.Printf("get from etcd failed,err:%v\n", err)
		return
	}
	for _, ev := range gr.Kvs {
		fmt.Printf("key:%s value:%s\n", ev.Key, ev.Value)
	}
	cancel()
}

func main() {
	//bodymain()
	test()
}
