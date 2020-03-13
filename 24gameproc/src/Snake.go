package main

import (
	"code.oldbody.com/studygolang/mylearn/24gameproc/src/Clib"
	"time"
)

//Shake 蛇子类
type Shake struct {
	size int                   //长度
	dir  byte                  //方向 byte-->uint8
	pos  [WIDE * HIGH]Position //定义数组存储每一节蛇的坐标
}

//初始化蛇信息
func (s *Shake) ShakeInit() {

	//初始化地图
	MapInit()
	//随机食物
	RandomFood()

	//蛇的长度
	s.size = 2
	//蛇头位置 从最中间开始的
	s.pos[0].X = WIDE / 2
	s.pos[0].Y = HIGH / 2
	//蛇尾位置
	s.pos[1].X = WIDE/2 - 1
	s.pos[1].Y = HIGH / 2
	//蛇的方向
	//用 U上 D下 L左 R右
	s.dir = 'R'
	//绘制蛇的UI
	for i := 0; i < s.size; i++ {
		var ch byte
		//区分蛇头跟蛇身
		if i == 0 {
			ch = '@'
		} else {
			ch = '#'
		}
		ShowUI(s.pos[i].X, s.pos[i].Y, ch)
	}

	//go 添加一个独立的函数，非阻塞执行
	//接收键盘中的输入
	go func() {
		for {
			switch Clib.Direction() {
			//根据键盘输入设置方向
			//方向上 W|w|向上的箭头
			case 72, 87, 119:
				if s.dir != 'D' {
					s.dir = 'U'
				}
			//方向左
			case 65, 97, 75:
				if s.dir != 'R' {
					s.dir = 'L'
				}
			//方向右
			case 100, 68, 77:
				if s.dir != 'L' {
					s.dir = 'R'
				}
			//方向下
			case 83, 115, 80:
				if s.dir != 'U' {
					s.dir = 'D'
				}
			//空格 暂停
			case 32:
				s.dir = 'P'
			}
		}
	}()
}

//开始游戏
func (s *Shake) PlayGame() {
	var dx, dy int = 0, 0
	//游戏流程控制，死循环
	for {
	FLAG:
		//延迟执行
		time.Sleep(time.Second / 6)
		//如果是空格，就暂停，直到s.dir接收到其他值
		if s.dir == 'P' {
			goto FLAG
		}
		//判断蛇的方向，更新蛇的位置
		switch s.dir {
		case 'U':
			dx = 0
			dy = -1
		case 'L':
			dx = -1
			dy = 0
		case 'R':
			dx = 1
			dy = 0
		case 'D':
			dx = 0
			dy = 1
		}
		//蛇头和墙碰撞
		if s.pos[0].X < 0+1 || s.pos[0].X >= WIDE || s.pos[0].Y < 0+3 || s.pos[0].Y >= HIGH+2 {
			return
		}
		//蛇头和蛇身碰撞
		for i := 1; i < s.size; i++ {
			if s.pos[0].X == s.pos[i].X && s.pos[0].Y == s.pos[i].Y {
				return
			}
		}
		//蛇头和食物碰撞
		if s.pos[0].X == food.X && s.pos[0].Y == food.Y {
			//身体增长
			s.size++
			//随机新食物
			RandomFood()
			//分数增加？？
			score++
			//根据分数调整速度
		}
		//记录蛇尾坐标
		lx := s.pos[s.size-1].X
		ly := s.pos[s.size-1].Y
		//更新蛇的坐标  蛇身坐标
		for i := s.size - 1; i > 0; i-- {
			s.pos[i].X = s.pos[i-1].X
			s.pos[i].Y = s.pos[i-1].Y
		}
		//蛇头坐标
		s.pos[0].X += dx
		s.pos[0].Y += dy
		//绘制蛇的UI
		for i := 0; i < s.size; i++ {
			var ch byte
			//区分蛇头跟蛇身
			if i == 0 {
				ch = '@'
			} else {
				ch = '*'
			}
			ShowUI(s.pos[i].X, s.pos[i].Y, ch)
		}
		//将蛇尾置空
		ShowUI(lx, ly, ' ')
	}
}
