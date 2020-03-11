package common

import (
	"fmt"
	"net"
	"strings"
)

//CollectEntry 要收集的日志的配置项的结构体
type CollectEntry struct {
	Path  string `json:"path"`  //去哪个日志读取日志文件
	Topic string `json:"topic"` //日志文件发往kafka中的哪个topic
}

//获取本机ip的函数
func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())                   //eg:（ip:端口号）
	ip = strings.Split(localAddr.IP.String(), ":")[0] //返回去掉端口号的ip地址
	return
}
