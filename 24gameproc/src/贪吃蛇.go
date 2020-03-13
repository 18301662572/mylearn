package main

import (
	"code.oldbody.com/studygolang/mylearn/24gameproc/src/Clib"
	"math/rand"
	"time"
)

func main() {
	//设置随机数种子，用作混淆操作
	rand.Seed(time.Now().UnixNano())
	//隐藏控制台光标
	Clib.HideCursor()

	//创建蛇对象
	s := new(Shake)
	//初始化蛇信息
	s.ShakeInit()
	//开始游戏
	s.PlayGame()
	//打印分数
	PrintScore()
}
