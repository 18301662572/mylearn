package main

import (
	"github.com/shirou/gopsutil/disk"
)

const (
	CpuInfoType  = "cpu"
	MemInfoType  = "mem"
	DiskInfoType = "disk"
	NetInfoType  = "net"
)

type MemInfo struct {
	Toltal      uint64  `json:"toltal"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type CpuInfo struct {
	CpuPercent float64 `json:"cpu_percent"`
}

type UsageStat struct {
	Path        string  `json:"path"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedpercent"`
}

type DiskInfo struct {
	PartitionUsageStat map[string]*disk.UsageStat
}

type IOStat struct {
	BytesSent       uint64
	BytesRecv       uint64
	PacketsSent     uint64
	PacketsRecv     uint64
	BytesSentRate   float64 `json:"bytesSentRate"`   // number of bytes sent
	BytesRecvRate   float64 `json:"bytesRecvRate"`   // number of bytes received
	PacketsSentRate float64 `json:"packetsSentRate"` // number of packets sent
	PacketsRecvRate float64 `json:"packetsRecvRate"` // number of packets received
}

type NetIOInfo struct {
	NetIOCountersStat map[string]*IOStat
}
