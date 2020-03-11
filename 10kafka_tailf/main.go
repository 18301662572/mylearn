package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

//tailf包打开文件，一直一条一条循环读取数据

func main() {
	filename := `./xx.log`
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	//1. 打开文件，获取数据
	tails, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Printf("tail %s failed,err:%v\n", filename, err)
		return
	}

	//2.循环从通道中读取数据
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines //chan tail.Line
		if !ok {
			fmt.Printf("tail file close reopen,filename:%s\n", tails.Filename)
			time.Sleep(time.Second) //读取出错，等一秒
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
