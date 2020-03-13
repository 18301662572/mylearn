# 贪吃蛇游戏项目

项目架构
```text
bin:编译程序
images:项目文档
src:代码，不参与程序编译
```

images包括
```text
1.模块设计
2.模块描述
3.项目分析
```

程序流程
```text
/*
type Position struct{
	X int
	Y int
}
1.蛇结构体创建
	长度
	方向
	[]坐标
2.食物的结构体创建
3.蛇初始化
	长度 2
	坐标
	方向
  地图初始化--调用C语言
  食物初始化
4.游戏流程控制
  蛇和墙碰撞 （蛇头）
  蛇和自身碰撞 （蛇头）
	蛇死亡
  蛇和食物碰撞 （蛇头）
	随机新食物
	蛇长度和身体增长
	分数 控制关卡和速度
5.游戏结束
	打印分数
	继续下关
*/
```