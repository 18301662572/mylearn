package main

import (
	"fmt"
	"github.com/sony/sonyflake"
)

//雪花算法
//id 生成服务

var (
	snoyFlake *sonyflake.Sonyflake
	machineID uint16 //真正的分布式环境下必须从zookeeper 或者etcd中获取Consul
)

//获取机器ID 的回调函数
func getMachineID() (uint16, error) {
	return machineID, nil
}

//Init初始化sonyFlake
func Init(mID uint16) (err error) {
	machineID = mID
	//配置项
	st := sonyflake.Settings{}
	st.MachineID = getMachineID
	snoyFlake = sonyflake.NewSonyflake(st)
	return
}

//GetID 获取全局唯一ID
func GetID() (id uint64, err error) {
	if snoyFlake == nil {
		err = fmt.Errorf("must Call Init before GetID,err:%v\n", err)
		return
	}
	return snoyFlake.NextID()
}

func main() {
	Init(1234567890)
	id, err := GetID()
	if err != nil {
		fmt.Printf("use sonyflake get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("use sonyflake get id success,ID:%d\n", id)
}
