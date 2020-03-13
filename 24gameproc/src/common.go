package main

import (
	"code.oldbody.com/studygolang/mylearn/24gameproc/src/Clib"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//全局常量 界面大小
const WIDE int = 20 //40
const HIGH int = 20

//全局变量 食物
var food Food

//全局分数
var score int = 0

//初始化父类 坐标
type Position struct {
	X int
	Y int
}

//Food 食物子类
type Food struct {
	Position
}

//初始化食物
func RandomFood() {
	//随机食物
	food.X = rand.Intn(WIDE) + 1 //0-19
	food.Y = rand.Intn(HIGH)

	//蛇跟食物重合之后重新随机（需要判断是否跟蛇有重合？？）
	//for i := 0; i < s.size; i++ {
	//	if food.X == s.pos[i].X && food.Y == s.pos[i].Y {
	//		//重新随机
	//	}
	//}

	//显示食物位置
	ShowUI(food.X, food.Y, '#')
}

//初始化地图
func MapInit() {
	//输出初始画面
	fmt.Fprintln(os.Stderr, //os.Stderr控制台接受者
		`
  #________________________________________#
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  |                                        |
  #________________________________________#
`)
}

//显示界面信息
func ShowUI(X, Y int, ch byte) {
	//调用C语言代码设置控制台光标
	//根据地图坐标有偏移
	Clib.GotoPosition(X*2+2, Y+2)
	//将字符绘制在界面中
	fmt.Fprintf(os.Stderr, "%c", ch)
}

//打印分数
func PrintScore() {
	//设置控制台光标位置
	Clib.GotoPosition(0, 23)
	fmt.Fprintln(os.Stderr, "分数：", score)
	time.Sleep(time.Second * 2)
}
