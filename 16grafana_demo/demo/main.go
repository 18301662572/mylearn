package main

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/shirou/gopsutil/cpu"
	"log"
	"time"
)

var cli client.Client

func initConnInflux() (err error) {
	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://192.168.42.133:8086",
		Username: "admin",
		Password: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	return
}

// insert
func writesPoints(percent float64) {
	//连接数据库
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "monitor", //数据库 monitor：监控
		Precision: "s",       //精度，默认ns
	})
	if err != nil {
		log.Fatalf("connect to test db failed,err:%v", err)
	}

	tags := map[string]string{"cpu": "cpu0"}
	fields := map[string]interface{}{
		"cpu_percent": percent,
	}
	//给表添加数据
	pt, err := client.NewPoint("cpu_percent", tags, fields, time.Now())
	if err != nil {
		log.Fatalf("client NewPoint failed,err:%v", err)
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		log.Fatalf("client write failed,err:%v", err)
	}
	log.Println("insert success")
}

func getCpuInfo() {
	// CPU使用率
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu percent:%v\n", percent)
	//写入到influxdb
	writesPoints(percent[0])
}

func main() {
	err := initConnInflux()
	if err != nil {
		fmt.Printf("connect to influxdb failed,err:%v\n", err)
		return
	}
	// 每一秒 getCpuInfo
	for {
		getCpuInfo()
		time.Sleep(time.Second)
	}
}
