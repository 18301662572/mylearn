package main

import "encoding/json"

//将连接器对象初始化
var h=hub{
	//connections注册了连接器
	connections:make(map[*connection]bool),
	//从连接器发送的信息
	broadcast:make(chan []byte),
	//从连接器注册请求
	register:make(chan *connection),
	//销毁请求
	unregister:make(chan *connection),
}


//处理ws的逻辑实现
func(h *hub) run(){
	//监听数据管道，在后端不断处理管道数据
	for{
		//根据不同的数据管道处理不同的逻辑
		select {
			//注册
			case c:=<-h.register:
				//标识注册了
				h.connections[c]=true
				//组装data数据
				c.data.Ip=c.ws.RemoteAddr().String()
				//更新类型
				c.data.Type="handshake"
				//用户列表
				c.data.UserList=user_list
				data_b, _ := json.Marshal(c.data)
				//将数据放入数据管道
				c.send<-data_b
			//注销
			case c:=<-h.register:
				//判断map里有要删除的数据
				if _,ok := h.connections[c];ok{
					delete(h.connections,c)
					close(c.send)
				}
			//连接中的信息
			case data:=<-h.broadcast:
				//处理数据流转，将数据同步到所有用户
				// c 是具体的每一个连接
				for c:= range h.connections{
					//将数据同步
					select {
					case c.send<-data:
					default:
						//防止死循环
						delete(h.connections,c)
						close(c.send )
					}
				}
		}
	}
}