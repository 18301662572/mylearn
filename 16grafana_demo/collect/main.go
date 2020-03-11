package main

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"log"
	"time"
)

var (
	cli                    client.Client
	lastNetIOStatTimeStamp int64      //上一次获取网络IO的时间点
	lastNetInfo            *NetIOInfo //上一次的网络IO数据
)

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

func getCpuInfo() {
	var cpuInfo = new(CpuInfo)
	// CPU使用率
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu percent:%v\n", percent)
	//写入到influxdb
	cpuInfo.CpuPercent = percent[0]
	writesCpuPoints(cpuInfo)
}

func getMemInfo() {
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed,err:%v", err)
		return
	}
	var memInfo = new(MemInfo)
	memInfo.Toltal = info.Total
	memInfo.Used = info.Used
	memInfo.UsedPercent = info.UsedPercent
	//写入到influxdb
	writesMemPoints(memInfo)
}

func getDiskInfo() {
	var diskInfo = &DiskInfo{
		PartitionUsageStat: make(map[string]*disk.UsageStat, 16),
	}
	parts, _ := disk.Partitions(true)
	for _, part := range parts {
		//拿到每一个分区
		usageStat, _ := disk.Usage(part.Mountpoint) //传挂载点
		diskInfo.PartitionUsageStat[part.Mountpoint] = usageStat
	}
	//写入到influxdb
	writesDiskPoints(diskInfo)
}

//net IO 网卡
func getNetInfo() {
	var netInfo = &NetIOInfo{
		NetIOCountersStat: make(map[string]*IOStat, 8),
	}
	currentTimeStamp := time.Now().Unix()
	netIOs, _ := net.IOCounters(true)
	for _, netIO := range netIOs {
		//fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
		var IoStat = new(IOStat)
		IoStat.BytesSent = netIO.BytesSent
		IoStat.BytesRecv = netIO.BytesRecv
		IoStat.PacketsRecv = netIO.PacketsRecv
		IoStat.PacketsSent = netIO.PacketsSent
		//将具体网卡数据的IoStat 变量添加到map
		netInfo.NetIOCountersStat[netIO.Name] = IoStat //不要放到contine下面
		//开始计算网卡相关速率
		if lastNetIOStatTimeStamp == 0 || lastNetInfo == nil {
			continue
		}
		//计算时间间隔
		interval := currentTimeStamp - lastNetIOStatTimeStamp
		IoStat.BytesSentRate = (float64(IoStat.BytesSent) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].BytesSent)) / float64(interval)
		IoStat.BytesRecvRate = (float64(IoStat.BytesRecv) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].BytesRecv)) / float64(interval)
		IoStat.PacketsSentRate = (float64(IoStat.PacketsSent) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].PacketsSent)) / float64(interval)
		IoStat.PacketsRecvRate = (float64(IoStat.PacketsRecv) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].PacketsRecv)) / float64(interval)
		//将具体网卡数据的IoStat变量添加到map中
		netInfo.NetIOCountersStat[netIO.Name] = IoStat
	}
	//更新全局记录的上一次采集网卡数据的时间点和网卡数据
	lastNetIOStatTimeStamp = currentTimeStamp
	lastNetInfo = netInfo
	//写入到influxdb
	writesNetIOPoints(netInfo)
}

func run(interval time.Duration) {
	ticker := time.Tick(interval)
	for _ = range ticker {
		getCpuInfo()
		getMemInfo()
		getDiskInfo()
		getNetInfo()
	}
}

func main() {
	err := initConnInflux()
	if err != nil {
		fmt.Printf("connect to influxdb failed,err:%v\n", err)
		return
	}
	// 每一秒 getCpuInfo
	run(time.Second)
}
