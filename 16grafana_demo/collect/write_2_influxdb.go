package main

import (
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)

// cpu
func writesCpuPoints(data *CpuInfo) {
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
		"cpu_percent": data.CpuPercent,
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
	log.Println("insert cpu success")
}

//mem
func writesMemPoints(data *MemInfo) {
	//连接数据库
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "monitor", //数据库 monitor：监控
		Precision: "s",       //精度，默认ns
	})
	if err != nil {
		log.Fatalf("connect to monitor db failed,err:%v", err)
	}
	tags := map[string]string{"mem": "mem"}
	fields := map[string]interface{}{
		"toltal":      int64(data.Toltal),
		"used":        int64(data.Used),
		"usedPercent": data.UsedPercent,
	}
	//给表添加数据
	pt, err := client.NewPoint("memory", tags, fields, time.Now())
	if err != nil {
		log.Fatalf("client NewPoint failed,err:%v", err)
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		log.Fatalf("client write failed,err:%v", err)
	}
	log.Println("insert mem success")
}

//disk
func writesDiskPoints(data *DiskInfo) {
	//连接数据库
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "monitor", //数据库 monitor：监控
		Precision: "s",       //精度，默认ns
	})
	if err != nil {
		log.Fatalf("connect to monitor db failed,err:%v", err)
	}
	//根据传入的数据的类型插入数据
	for k, v := range data.PartitionUsageStat {
		tags := map[string]string{"path": k}
		fields := map[string]interface{}{
			"toltal":      int64(v.Total),
			"free":        int64(v.Free),
			"used":        int64(v.Used),
			"usedPercent": v.UsedPercent,
		}
		//给表添加数据
		pt, err := client.NewPoint("disk", tags, fields, time.Now())
		if err != nil {
			log.Fatalf("client NewPoint failed,err:%v", err)
		}
		bp.AddPoint(pt)
	}
	err = cli.Write(bp)
	if err != nil {
		log.Fatalf("client write failed,err:%v", err)
	}
	log.Println("insert disk success")
}

//net io 网卡
func writesNetIOPoints(data *NetIOInfo) {
	//连接数据库
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "monitor", //数据库 monitor：监控
		Precision: "s",       //精度，默认ns
	})
	if err != nil {
		log.Fatalf("connect to test db failed,err:%v", err)
	}
	//根据传入的数据的类型插入数据
	for k, v := range data.NetIOCountersStat {
		tags := map[string]string{"name": k} //每一个网卡存为tag
		fields := map[string]interface{}{
			"bytesSentRate":   v.BytesSentRate,
			"bytesRecvRate":   v.BytesRecvRate,
			"packetsSentRate": v.PacketsSentRate,
			"packetsRecvRate": v.PacketsRecvRate,
		}
		//给表添加数据
		pt, err := client.NewPoint("netio", tags, fields, time.Now())
		if err != nil {
			log.Fatalf("client NewPoint failed,err:%v", err)
		}
		bp.AddPoint(pt)
	}
	err = cli.Write(bp)
	if err != nil {
		log.Fatalf("client write failed,err:%v", err)
	}
	log.Println("insert netio success")
}
